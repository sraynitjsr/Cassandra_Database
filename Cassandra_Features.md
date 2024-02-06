## What is Cassandra?

#### Cassandra is a distributed NoSQL database system that was originally developed by Facebook and later open-sourced as an Apache project.
#### It is designed to handle large amounts of data spread out across many commodity servers while providing high availability and fault tolerance.
#### Cassandra is optimized for write-heavy workloads and offers linear scalability by adding more nodes to the cluster.
#### Cassandra is a highly scalable distributed NoSQL database.
#### Cassandra is designed to handle large amounts of data across many commodity servers while providing high availability and fault tolerance.

## Key Concepts:-

#### a. Nodes: Cassandra is a distributed database, and data is distributed across multiple nodes. Each node can serve read and write requests independently.

#### b. Clusters: A Cassandra cluster is a group of nodes that work together. Data is replicated across multiple nodes in the cluster for fault tolerance.

#### c. Keyspaces: Keyspaces are the top-level containers for data in Cassandra, similar to databases in traditional relational databases. Each keyspace contains multiple tables.

#### d. Tables: Tables in Cassandra are similar to tables in relational databases but have a more flexible schema. Each table has a primary key, which uniquely identifies rows in the table.

#### e. Partitioning: Cassandra distributes data across multiple nodes using partitioning. Each row in a table is identified by a partition key, and Cassandra uses a partitioner to determine which node in the cluster should store the data based on the partition key.

#### f. Replication: Cassandra replicates data across multiple nodes in the cluster to ensure high availability and fault tolerance. You can configure the replication strategy and the number of replicas for each keyspace.

#### g. Consistency Levels: Cassandra offers tunable consistency levels, allowing you to trade off consistency for availability and performance. Consistency levels determine how many replicas must respond to a read or write operation for it to be considered successful.
