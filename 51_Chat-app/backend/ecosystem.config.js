module.exports = {
  apps: [{
    name: "chat-app",
    script: "./chat-app",
    watch: true,
    env: {
      "PORT": "8080",
      "NODE_ENV": "production",
    },
    max_memory_restart: "1G",
    instances: 1,
    autorestart: true,
    restart_delay: 4000,
    error_file: "logs/err.log",
    out_file: "logs/out.log",
    log_file: "logs/combined.log",
    time: true
  }]
} 