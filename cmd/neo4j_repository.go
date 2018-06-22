package cmd

import "github.com/johnnadratowski/golang-neo4j-bolt-driver"

func createTeamEdge(conn golangNeo4jBoltDriver.Conn, t *team) error {
	_, err := conn.ExecNeo(
		"CREATE (:Team {teamID: {teamID}, name: {name}, createAt: {createAt}, deleteAt: {deleteAt}, type: 'open'})",
		map[string]interface{}{"teamID": t.id, "name": t.name, "createAt": t.createAt, "deleteAt": t.deleteAt, "type": t.displayType()},
	)
	if err != nil {
		return err
	}
	return nil
}

func createChannelEdge(conn golangNeo4jBoltDriver.Conn, c *channel) error {
	_, err := conn.ExecNeo(
		"CREATE (:Channel {channelID: {channelID}, name: {name}, createAt: {createAt}, deleteAt: {deleteAt}, type: 'open'})",
		map[string]interface{}{"channelID": c.id, "name": c.name, "createAt": c.createAt, "deleteAt": c.deleteAt, "type": c.displayType()},
	)
	if err != nil {
		return err
	}
	return nil
}

func createUserEdge(conn golangNeo4jBoltDriver.Conn, u *user) error {
	_, err := conn.ExecNeo(
		"CREATE (:User {userID: {userID}, username: {username}, createAt: {createAt}, deleteAt: {deleteAt}})",
		map[string]interface{}{"userID": u.id, "username": u.username, "createAt": u.createAt, "deleteAt": u.deleteAt},
	)
	if err != nil {
		return err
	}
	return nil
}

func createPostEdge(conn golangNeo4jBoltDriver.Conn, p *post) error {
	_, err := conn.ExecNeo(
		"CREATE (:Post {postID: {postID}, message: {message}, createAt: {createAt}, deleteAt: {deleteAt}})",
		map[string]interface{}{"postID": p.id, "message": p.message, "createAt": p.createAt, "deleteAt": p.deleteAt},
	)
	if err != nil {
		return err
	}
	return nil
}
