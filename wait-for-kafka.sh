#!/bin/sh

CONTAINER_NAME="broker"
TIMEOUT=60

echo "Waiting for Kafka container ($CONTAINER_NAME) to be healthy..."

for i in $(seq 1 ${TIMEOUT}); do
  if docker inspect --format="{{json .State.Health.Status}}" $CONTAINER_NAME | grep "healthy" > /dev/null; then
    echo "Kafka container ($CONTAINER_NAME) is now healthy."
    exit 0
  fi
  sleep 1
done

echo "Timed out waiting for Kafka container ($CONTAINER_NAME) to be healthy."
exit 1
