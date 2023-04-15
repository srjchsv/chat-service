## Chat-service for a real-time chat app

Build using microservices architecture

![Real-Time Chat System Design](/chat.png "Real-Time Chat System Design")

## Stack

[Frontend](https://github.com/srjchsv/chat-frontend): Next js

Backend:

- [auth-service](https://github.com/srjchsv/auth-service): go, gin-gonic, postgres, gorm, jwt
- [chat-service](https://github.com/srjchsv/chat-service): go, gin-gonic, postgres, gorm, jwt, confluent-kafka, gorilla websocket
- [notifications-service](https://github.com/srjchsv/notifications-service): go, echo, gorilla websocket, confluent-kafka

To run the backend services use `make run` command. It will run docker compose if needed and then go.

For frontend use `npm i` and `npm run dev`
<<<<<<< HEAD

To be done (optional):

- Built Docker containers images and push to registry
- Kubernetes deployment and service yamls
- Run on k8s cluster
=======
>>>>>>> 98be661 (initial commit)
