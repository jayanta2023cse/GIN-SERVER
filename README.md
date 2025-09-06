# Basic Gin Server with GORM Database Connection

This is a basic Golang web server built using the [Gin](https://github.com/gin-gonic/gin) web framework. It includes database connectivity using [GORM](https://gorm.io/), supporting MySQL, PostgreSQL, and MSSQL.

## 📦 Features

- 🚀 Lightweight HTTP server using Gin
- 🔌 Database connection using GORM
- 🔁 Retry logic for DB connection
- 📁 Configurable via `.env` or config file
- 🧪 Ready for REST API development

## 🛠️ Tech Stack

- Go 1.18+
- Gin Web Framework
- GORM ORM
- MySQL / PostgreSQL / SQL Server (via GORM drivers)
- Kafka

## 🔗 Server Running At

Your server will be accessible at: http://localhost:8080

# Why Use Kafka in a Single-Server Project?

Even in a monolith or single-server application, **Apache Kafka** can bring significant benefits:

- **Decoupling** → Producers and consumers don’t depend on each other’s runtime.
- **Replay** → Events are persisted and can be reprocessed anytime.
- **Scalability-ready** → If you later migrate to microservices, Kafka already fits in seamlessly.
- **Asynchronous processing** → Long-running tasks don’t block your main request cycle.

---

## 🔹 Example Use Cases in Single-Server Applications

### 1. Audit Logging

- Instead of writing logs directly to the database, publish events to Kafka.
- Consumers can process and store them (or forward them to ELK/Splunk for analysis).

---

### 2. Background Jobs

- Example: A user uploads a file.
- The producer publishes a `"file_uploaded"` event.
- A consumer picks it up and processes the file in the background.

---

### 3. Event-Driven Workflows

- Example: `"User registered"` event.
- Consumers can trigger:
  - Sending a welcome email
  - Creating a default profile
  - Logging the registration for analytics

Even if these run on the same server, they remain **loosely coupled**.

---

## ✅ Takeaway

Kafka isn’t just for microservices or big data.  
It can add **reliability, flexibility, and future scalability** even to a **single-server monolith**.

### ✅ Create Kafka test-topic:

<!-- In Local or Bash Terminal follow this format localhost:9093 -->

docker exec kafka kafka-topics --create --topic test-topic --bootstrap-server localhost:9093 --partitions 1 --replication-factor 1

<!-- Inside Docker Container follow this format kafka:9092 -->

kafka-topics --create --topic test-topic --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1

### ✅ Verify Current Topic List:

docker exec kafka kafka-topics --list --bootstrap-server localhost:9093

### ✅ Start the Produces:

docker exec -i kafka kafka-console-producer --broker-list localhost:9093 --topic test-topic

### ✅ Start the Consumer:

docker exec -i kafka kafka-console-consumer --bootstrap-server localhost:9093 --topic test-topic --from-beginning

### ✅ How to Delete a Topic:

docker exec kafka kafka-topics --bootstrap-server localhost:9093 --delete --topic <Topic-Name>

### ✅ Check Current Topic Configuration:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --describe --topic test-topic

<!-- Look for retention.ms (default is 604800000 ms, or 7 days). -->

### ✅ Set Low Retention Period:

<!-- Modify the topic’s retention period to a short duration (e.g., 1000 ms = 1 second) to delete all messages: -->

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config retention.ms=1000

### ✅ Restore Retention Period (Optional):

<!-- Reset the retention period to a reasonable value (e.g., 7 days): -->

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config retention.ms=604800000

### ✅ Reset Consumer Offsets (Optional):

<!-- If your consumer group (test-group) has offsets, reset them to start from the beginning of the topic: -->

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --group test-group --reset-offsets --to-earliest --topic test-topic --execute

<!-- Verify the reset: -->

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --group test-group --describe
