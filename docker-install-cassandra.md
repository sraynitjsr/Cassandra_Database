## Pull the Cassandra Docker Image
#### docker pull cassandra

## Run Cassandra Container
#### docker run --name my-cassandra-container -p 9042:9042 -d cassandra

## docker ps
#### Check for cassandra containers, for further steps

## Connect to Cassandra from CQLSH
#### docker exec -it my-cassandra-container cqlsh
