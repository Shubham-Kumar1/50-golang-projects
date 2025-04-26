import { useState, useEffect, useRef } from 'react';
import { Box, Container, TextField, Button, Typography, Paper, Stack, Divider } from '@mui/material';
import { wsService, Message } from '../services/websocket';

export const Chat = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [newMessage, setNewMessage] = useState('');
  const messagesEndRef = useRef<HTMLDivElement>(null);
  const isConnected = useRef(false);

  useEffect(() => {
    // Only connect if not already connected
    if (!isConnected.current) {
      console.log('Initializing WebSocket connection');
      wsService.connect();
      isConnected.current = true;
    }

    // Register message handler
    const unsubscribe = wsService.onMessage((message) => {
      console.log('Message received in Chat component:', message);
      
      // Only add system messages if they're not duplicates
      if (message.isSystemMessage) {
        setMessages((prev) => {
          // Check if the last message is the same system message
          const lastMessage = prev[prev.length - 1];
          if (lastMessage && 
              lastMessage.isSystemMessage && 
              lastMessage.body === message.body) {
            return prev; // Don't add duplicate system messages
          }
          return [...prev, message];
        });
      } else {
        // Check if the message body is a JSON string and parse it
        let processedMessage = { ...message };
        try {
          if (typeof message.body === 'string' && message.body.startsWith('{') && message.body.endsWith('}')) {
            const parsedBody = JSON.parse(message.body);
            if (parsedBody.body) {
              processedMessage.body = parsedBody.body;
            }
          }
        } catch (e) {
          // If parsing fails, use the original message
          console.error('Failed to parse message body:', e);
        }
        
        setMessages((prev) => [...prev, processedMessage]);
      }
    });

    // Cleanup function
    return () => {
      console.log('Cleaning up Chat component');
      unsubscribe();
      // Don't disconnect here, as it might be used by other components
      // wsService.disconnect();
      // isConnected.current = false;
    };
  }, []);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const handleSendMessage = (e: React.FormEvent) => {
    e.preventDefault();
    if (newMessage.trim()) {
      console.log('Sending message from Chat component:', newMessage);
      wsService.sendMessage(newMessage);
      setNewMessage('');
    }
  };

  return (
    <Container maxWidth="sm" sx={{ height: '100vh', py: 4 }}>
      <Paper elevation={3} sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
        <Box sx={{ p: 2, borderBottom: 1, borderColor: 'divider' }}>
          <Typography variant="h6">Chat Room</Typography>
        </Box>
        
        <Box sx={{ flex: 1, overflow: 'auto', p: 2 }}>
          <Stack spacing={2}>
            {messages.map((message, index) => (
              message.isSystemMessage ? (
                <Box key={index} sx={{ display: 'flex', justifyContent: 'center', my: 1 }}>
                  <Typography variant="caption" color="text.secondary" sx={{ 
                    bgcolor: 'grey.200', 
                    px: 1, 
                    py: 0.5, 
                    borderRadius: 1,
                    fontSize: '0.75rem'
                  }}>
                    {message.body}
                  </Typography>
                </Box>
              ) : (
                <Paper
                  key={index}
                  elevation={1}
                  sx={{
                    p: 1,
                    bgcolor: message.type === 1 ? 'primary.light' : 'grey.100',
                    alignSelf: message.type === 1 ? 'flex-end' : 'flex-start',
                    maxWidth: '80%',
                  }}
                >
                  <Typography variant="body1">{message.body}</Typography>
                  <Typography variant="caption" color="text.secondary">
                    {new Date(message.timestamp).toLocaleTimeString()}
                  </Typography>
                </Paper>
              )
            ))}
            <div ref={messagesEndRef} />
          </Stack>
        </Box>

        <Box component="form" onSubmit={handleSendMessage} sx={{ p: 2, borderTop: 1, borderColor: 'divider' }}>
          <Stack direction="row" spacing={1}>
            <TextField
              fullWidth
              size="small"
              value={newMessage}
              onChange={(e) => setNewMessage(e.target.value)}
              placeholder="Type a message..."
              variant="outlined"
            />
            <Button type="submit" variant="contained">
              Send
            </Button>
          </Stack>
        </Box>
      </Paper>
    </Container>
  );
}; 