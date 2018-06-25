# mattergraph
CLI to get Mattermost data into Neo4j

```bash
docker pull neo4j
docker run --publish=7474:7474 --publish=7687:7687 --volume=$HOME/neo4j/data:/data neo4j
open http://localhost:7474/
```
Username: neo4j\
Password: neo4j
```bash
cp .mattergraph.example.yaml $HOME/.mattergraph.yaml
go run main.go etl
```
