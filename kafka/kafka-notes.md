# Kafka Implementation Notes

<!-- Use localhost:9093 for local Development -->
<!-- Use localhost:9092 for Docker Production Development -->

### ✅ List All Topics:

docker exec kafka kafka-topics --list --bootstrap-server localhost:9093

### ✅ Describe a Topic:

docker exec kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:9093

### ✅ Alter Topic Partitions:

docker exec kafka kafka-topics --alter --topic test-topic --bootstrap-server localhost:9093 --partitions 3

### ✅ List Consumer Groups:

docker exec kafka kafka-consumer-groups --list --bootstrap-server localhost:9093

### ✅ Describe Consumer Group:

docker exec kafka kafka-consumer-groups --describe --group test-group --bootstrap-server localhost:9093

### ✅ Increase Topic Replication Factor:

docker exec kafka kafka-topics --alter --topic test-topic --bootstrap-server localhost:9093 --add-config min.insync.replicas=2

### ✅ Produce Messages from File:

cat messages.txt | docker exec -i kafka kafka-console-producer --broker-list localhost:9093 --topic test-topic

### ✅ Consume Messages with Group:

docker exec -i kafka kafka-console-consumer --bootstrap-server localhost:9093 --topic test-topic --group test-group --from-beginning

### ✅ Reset Offsets to Latest:

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --group test-group --reset-offsets --to-latest --topic test-topic --execute

### ✅ Set Message Max Size:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config max.message.bytes=2000000

### ✅ View Cluster Metadata:

docker exec kafka kafka-cluster cluster-metadata --bootstrap-server localhost:9093

### ✅ List ACLs:

docker exec kafka kafka-acls --list --bootstrap-server localhost:9093

### ✅ Add ACL for Consumer:

docker exec kafka kafka-acls --add --allow-principal User:consumer --operation Read --topic test-topic --bootstrap-server localhost:9093

### ✅ Get Broker Configs:

docker exec kafka kafka-configs --describe --bootstrap-server localhost:9093 --entity-type brokers --entity-name 1

### ✅ Reassign Partitions:

docker exec kafka kafka-reassign-partitions --bootstrap-server localhost:9093 --topics-to-move-json-file topics.json --broker-list 1,2 --generate

### ✅ Check Under-Replicated Partitions:

docker exec kafka kafka-topics --under-replicated-partitions --bootstrap-server localhost:9093

### ✅ Dump Log Segments:

docker exec kafka kafka-dump-log --files /var/lib/kafka/data/test-topic-0/00000000000000000000.log --print-data-log

### ✅ Get Topic Offsets:

docker exec kafka kafka-get-offsets --topic test-topic --bootstrap-server localhost:9093

### ✅ Run Preferred Leader Election:

docker exec kafka kafka-preferred-replica-election --bootstrap-server localhost:9093 --topics test-topic

### ✅ List Brokers:

docker exec kafka kafka-broker-api-versions --bootstrap-server localhost:9093

### ✅ Update Broker Config:

docker exec kafka kafka-configs --alter --bootstrap-server localhost:9093 --entity-type brokers --entity-name 1 --add-config log.cleaner.enable=true

### ✅ Monitor Topic Lag:

docker exec kafka kafka-consumer-groups --describe --group test-group --bootstrap-server localhost:9093 --offsets

### ✅ Create Snapshot of Offsets:

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --group test-group --export-offsets --output-file offsets.json

### ✅ Seek Consumer to Offset:

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --group test-group --reset-offsets --to-offset 100 --topic test-topic:0 --execute

### ✅ Verify Topic Data:

docker exec -i kafka kafka-console-consumer --bootstrap-server localhost:9093 --topic test-topic --max-messages 10 --from-beginning

### ✅ Set Compression Type:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config compression.type=gzip

### ✅ List All Configs for Topic:

docker exec kafka kafka-configs --describe --topic test-topic --bootstrap-server localhost:9093

### ✅ Remove Config from Topic:

docker exec kafka kafka-configs --alter --topic test-topic --delete-config retention.ms --bootstrap-server localhost:9093

### ✅ Check Kafka Version:

docker exec kafka kafka-broker-api-versions --bootstrap-server localhost:9093 | grep version

### ✅ Generate Mirror Maker Config:

docker exec kafka kafka-mirror-maker --new.consumer --bootstrap-server localhost:9093 --topics test-topic --generate-config

### ✅ Run Performance Test for Producer:

docker exec kafka kafka-producer-perf-test --topic test-topic --num-records 100000 --record-size 100 --throughput -1 --producer-props bootstrap.servers=localhost:9093

### ✅ Run Performance Test for Consumer:

docker exec kafka kafka-consumer-perf-test --broker-list localhost:9093 --topic test-topic --messages 100000

### ✅ Enable Debug Logging Temporarily:

docker exec kafka kafka-configs --alter --bootstrap-server localhost:9093 --entity-type brokers --entity-name 1 --add-config log4j.logger.kafka=DEBUG

### ✅ Disable Debug Logging:

docker exec kafka kafka-configs --alter --bootstrap-server localhost:9093 --entity-type brokers --entity-name 1 --delete-config log4j.logger.kafka

### ✅ List All Topics with Partitions:

docker exec kafka kafka-topics --list --bootstrap-server localhost:9093 | xargs -I {} docker exec kafka kafka-topics --describe --topic {} --bootstrap-server localhost:9093

### ✅ Backup Topic Data:

docker exec kafka kafka-dump-log --files /var/lib/kafka/data/test-topic-0/\_.log --print-data-log > backup.log

### ✅ Check Disk Usage for Topic:

docker exec kafka du -sh /var/lib/kafka/data/test-topic-\_

### ✅ Compact Topic (for Log Compaction):

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config cleanup.policy=compact

### ✅ Disable Log Compaction:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config cleanup.policy=delete

### ✅ Set Min In-Sync Replicas:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config min.insync.replicas=1

### ✅ Get Leader for Topic:

docker exec kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:9093 | grep Leader

### ✅ Force Leader Election:

docker exec kafka kafka-leader-election --bootstrap-server localhost:9093 --topic test-topic --partition 0 --election-type preferred

### ✅ View Consumer Lag:

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --group test-group --describe | awk '{lag+=$6} END {print lag}'

### ✅ Increase Topic Retention Size:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config retention.bytes=1073741824

### ✅ Check Broker Rack:

docker exec kafka kafka-broker-api-versions --bootstrap-server localhost:9093 | grep rack

### ✅ Add Partition to Topic:

docker exec kafka kafka-topics --alter --topic test-topic --partitions 4 --bootstrap-server localhost:9093

### ✅ Describe Cluster:

docker exec kafka kafka-cluster cluster-metadata --bootstrap-server localhost:9093

### ✅ Test SSL Connection (if enabled):

docker exec kafka kafka-console-producer --broker-list localhost:9093 --topic test-topic --producer.config client-ssl.properties

### ✅ Generate Client Config for SASL (if enabled):

docker exec kafka kafka-configs --describe --entity-type users --entity-name client --bootstrap-server localhost:9093

### ✅ Run Verifiable Producer:

docker exec kafka kafka-verifiable-producer --broker-list localhost:9093 --topic test-topic --max-messages 10

### ✅ Run Verifiable Consumer:

docker exec kafka kafka-verifiable-consumer --broker-list localhost:9093 --topic test-topic --group-id test-group

### ✅ List All Broker Logs:

docker exec kafka ls /var/log/kafka

### ✅ Tail Kafka Server Log:

docker exec -it kafka tail -f /var/log/kafka/server.log

### ✅ Check Kafka Health:

docker exec kafka kafka-broker-api-versions --bootstrap-server localhost:9093 > /dev/null && echo "Kafka is healthy" || echo "Kafka is unhealthy"

### ✅ Set Segment Size:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config segment.bytes=107374182

### ✅ Get End Offsets:

docker exec kafka kafka-get-offsets --topic test-topic --time -1 --bootstrap-server localhost:9093

### ✅ Get Beginning Offsets:

docker exec kafka kafka-get-offsets --topic test-topic --time -2 --bootstrap-server localhost:9093

### ✅ Delete Consumer Group:

docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9093 --delete --group test-group

### ✅ List Deleted Topics:

docker exec kafka kafka-topics --list --bootstrap-server localhost:9093 --exclude-internal | grep \_\_deleted

### ✅ Recover from Unclean Shutdown:

docker exec kafka kafka-server-start /etc/kafka/server.properties --override recovery.threads.per.data.dir=1

### ✅ Set Flush Interval:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --topic test-topic --add-config flush.ms=1000

### ✅ Monitor Broker Metrics:

docker exec kafka jmxtool -query kafka.server:type=BrokerTopicMetrics,name=TotalProduceRequestsPerSec

### ✅ Enable JMX for Monitoring:

docker run ... -e JMX_PORT=9999 ...

### ✅ Check Controller Status:

docker exec kafka kafka-controller-quorum-voters --bootstrap-server localhost:9093 --describe

### ✅ Scale Brokers (Manual Step):

### ✅ Encrypt Data at Rest (Advanced):

### ✅ Set ACL for Producer:

docker exec kafka kafka-acls --add --allow-principal User:producer --operation Write --topic test-topic --bootstrap-server localhost:9093

### ✅ Remove ACL:

docker exec kafka kafka-acls --remove --allow-principal User:producer --operation Write --topic test-topic --bootstrap-server localhost:9093

### ✅ List Internal Topics:

docker exec kafka kafka-topics --list --bootstrap-server localhost:9093 --exclude-internal=false | grep \*\*

### ✅ Compact Internal Topic:

docker exec kafka kafka-configs --bootstrap-server localhost:9093 --alter --entity-type topics --entity-name \*\*consumer_offsets --add-config cleanup.policy=compact

### ✅ View Log Cleaner Status:

docker exec kafka jmxtool -query kafka.log:type=LogCleaner,name=LogCleanerManager

### ✅ Set Log Cleaner Threads:

docker exec kafka kafka-configs --alter --bootstrap-server localhost:9093 --entity-type brokers --entity-name 1 --add-config log.cleaner.threads=2

### ✅ Get Replica Lag:

docker exec kafka kafka-replica-unfollowers --bootstrap-server localhost:9093

### ✅ Run ISR Shrink/Expand Test:

docker exec kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:9093 | grep Isr

### ✅ Configure Rack Awareness:

docker exec kafka kafka-configs --alter --bootstrap-server localhost:9093 --entity-type brokers --entity-name 1 --add-config rack=zone1

### ✅ Verify Rack Awareness:

docker exec kafka kafka-topics --describe --topic test-topic --bootstrap-server localhost:9093
