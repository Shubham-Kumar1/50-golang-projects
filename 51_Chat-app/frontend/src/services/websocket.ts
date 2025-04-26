export interface Message {
  type: number;
  body: string;
  timestamp: string;
  client_id?: string;
  isSystemMessage?: boolean;
}

class WebSocketService {
  private ws: WebSocket | null = null;
  private messageHandlers: ((message: Message) => void)[] = [];
  private clientId: string | null = null;
  private isConnecting: boolean = false;
  private reconnectTimeout: number | null = null;

  connect() {
    // Prevent multiple simultaneous connection attempts
    if (this.isConnecting || this.ws?.readyState === WebSocket.OPEN) {
      console.log('WebSocket already connected or connecting');
      return;
    }

    this.isConnecting = true;
    console.log('Connecting to WebSocket...');
    
    this.ws = new WebSocket('ws://localhost:9090/ws');

    this.ws.onopen = () => {
      console.log('Connected to WebSocket');
      this.isConnecting = false;
      
      // Clear any pending reconnect timeout
      if (this.reconnectTimeout !== null) {
        clearTimeout(this.reconnectTimeout);
        this.reconnectTimeout = null;
      }
    };

    this.ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        console.log('Received message:', data);
        
        // Check if this is a system message
        const isSystemMessage = data.body === "New User Added" || data.body === "One User Disconnected";
        
        // Ensure timestamp is a string
        const timestamp = data.timestamp || new Date().toISOString();
        
        const message: Message = {
          ...data,
          isSystemMessage,
          timestamp: typeof timestamp === 'string' ? timestamp : new Date(timestamp).toISOString()
        };
        
        this.messageHandlers.forEach(handler => handler(message));
      } catch (error) {
        console.error('Error parsing WebSocket message:', error);
      }
    };

    this.ws.onclose = () => {
      console.log('Disconnected from WebSocket');
      this.isConnecting = false;
      
      // Attempt to reconnect after 5 seconds
      this.reconnectTimeout = window.setTimeout(() => this.connect(), 5000);
    };

    this.ws.onerror = (error) => {
      console.error('WebSocket error:', error);
      this.isConnecting = false;
    };
  }

  sendMessage(message: string) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      console.log('Sending message:', message);
      this.ws.send(JSON.stringify({
        type: 1,
        body: message
      }));
    } else {
      console.error('Cannot send message: WebSocket not connected');
    }
  }

  onMessage(handler: (message: Message) => void) {
    console.log('Registering message handler');
    this.messageHandlers.push(handler);
    return () => {
      console.log('Unregistering message handler');
      this.messageHandlers = this.messageHandlers.filter(h => h !== handler);
    };
  }

  disconnect() {
    console.log('Disconnecting WebSocket');
    if (this.reconnectTimeout !== null) {
      clearTimeout(this.reconnectTimeout);
      this.reconnectTimeout = null;
    }
    
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
    
    this.isConnecting = false;
  }
}

export const wsService = new WebSocketService(); 