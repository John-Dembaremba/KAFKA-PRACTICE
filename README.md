# KAFKA-PRAC

## Overview
KAFKA-PRAC is a project aimed at practicing Apache Kafka implementation with a single-node setup. It utilizes Kafka for message streaming and Kafka-UI for observability. Producers and consumers are written in Go language for message production and consumption.

## Prerequisites
Before running the project, ensure the following prerequisites are met:
- Docker
- Python
- Go

## Getting Started
Follow these steps to set up and run the KAFKA-PRAC project:

1. Run `python3 client_id.py`. This will generate a UUID for the cluster ID. Place this UUID in the `CLUSTER_ID` field in the Docker Compose file.

2. Run `docker-compose up` to create the Kafka and Kafka-UI services.

3. Execute `go run main.go` to start producing and consuming messages.

4. Open Kafka-UI at [http://localhost:8080/](http://localhost:8080/) to observe Kafka metrics and activity.

## Components
- **Apache Kafka**: A distributed streaming platform used for building real-time data pipelines and streaming applications.
- **Kafka-UI**: Provides a user interface for monitoring and managing Kafka clusters.
- **Go (Golang)**: Used for writing producers and consumers for Kafka messages.

## Notes
- The project assumes a single-node Kafka setup for simplicity.
- Producers and consumers are written in Go to demonstrate message production and consumption.
- Kafka-UI facilitates monitoring and observability of Kafka cluster activity.

