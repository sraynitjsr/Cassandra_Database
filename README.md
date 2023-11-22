## Cassandra is a highly scalable and distributed NoSQL database system designed for handling large amounts of data across multiple commodity servers without any single point of failure.

#### Distributed and Decentralized: Cassandra is designed to be distributed, allowing it to run on multiple nodes across different locations or data centers. It follows a decentralized architecture, and there is no single point of failure.

#### Scalability: Cassandra provides linear scalability, meaning we can add more nodes to the cluster to handle increased data and traffic. This makes it well-suited for applications with high read and write throughput requirements.

#### No Single Point of Failure: Data in Cassandra is replicated across multiple nodes in the cluster, ensuring that there is no single point of failure. This replication strategy contributes to both fault tolerance and high availability.

#### Column-Family Data Model: Cassandra organizes data into column families, which are similar to tables in a relational database. Each row in a column family is identified by a unique key, and each row can have a dynamic number of columns.

#### Query Language (CQL): Cassandra Query Language (CQL) is a SQL-like language for interacting with Cassandra. It provides a familiar syntax for developers who are already familiar with SQL databases, making it easier to transition to Cassandra.

#### Tunable Consistency: Cassandra provides tunable consistency, allowing developers to choose the level of consistency they need for each read and write operation. This flexibility is essential for balancing performance and consistency in distributed systems.

#### Write and Read Performance: Cassandra is optimized for high write and read throughput. It supports both fast writes and efficient reads, making it suitable for use cases with heavy write loads and large-scale data retrieval.

#### Support for Time Series Data: Cassandra is well-suited for handling time series data due to its ability to efficiently write and retrieve data based on timestamps.

#### Wide-Column Store: Cassandra is often classified as a wide-column store NoSQL database. This means that data is stored in columns rather than rows, allowing for efficient querying and retrieval of specific columns of data.

## 1.Understanding the Data Model:-
#### Cassandra uses a NoSQL data model, specifically a wide-column store. The basic building blocks are:
#### Keyspace: The top-level container for data in Cassandra.
#### Column Family (Table): Represents a collection of rows, each with a unique key. Each row consists of columns, and columns are grouped into column families.
#### Column: The basic unit of data in Cassandra. A column has a name, a value, and a timestamp. Columns are grouped into rows.

## 2. Basic CQL Commands:-
#### CREATE KEYSPACE my_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
#### USE my_keyspace;
#### CREATE TABLE my_table (id UUID PRIMARY KEY, name TEXT, age INT);
#### INSERT INTO my_table (id, name, age) VALUES (uuid(), 'John Doe', 25);

