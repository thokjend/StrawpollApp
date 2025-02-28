export class WebSocketService {
  private socket: WebSocket | null = null;
  private onMessageCallback: ((event: MessageEvent) => void) | null = null;

  constructor(private url: string) {}

  connect() {
    this.socket = new WebSocket(this.url);

    this.socket.onopen = () => {
      console.log("WebSocket connected");
    };

    this.socket.onmessage = (event) => {
      console.log("WebSocket message received:", event.data);
      if (this.onMessageCallback) {
        this.onMessageCallback(event);
      }
    };

    this.socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    this.socket.onclose = () => {
      console.log("WebSocket disconnected");
    };
  }

  setOnMessageCallback(callback: (event: MessageEvent) => void) {
    this.onMessageCallback = callback;
  }

  disconnect() {
    if (this.socket) {
      this.socket.close();
    }
  }
}
