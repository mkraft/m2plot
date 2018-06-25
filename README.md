# mattergraph
CLI to get Mattermost data into Neo4j

## Setup
```bash
docker pull neo4j
docker run --publish=7474:7474 --publish=7687:7687 --volume=$HOME/neo4j/data:/data neo4j
open http://localhost:7474/
```
Login with `neo4j` \ `neo4j` and reset the password.

```bash
cp .mattergraph.example.yaml $HOME/.mattergraph.yaml
```

Update the Neo4j connection string in `~/.mattergraph.yaml` with the new password.

## Usage

```bash
go run main.go etl
```
