// Copyright © 2018 Martin Kraft <martinkraft@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func createTeamMemberVertex(conn golangNeo4jBoltDriver.Conn, tm *teamMember) error {
	_, err := conn.ExecNeo(
		`MATCH
			(user:User {userID: {userID}}),
			(team:Team {teamID: {teamID}})
		CREATE (user)-[:MEMBER_OF]->(team)`,
		map[string]interface{}{"teamID": tm.teamID, "userID": tm.userID},
	)
	if err != nil {
		return err
	}
	return nil
}

func createChannelMemberVertex(conn golangNeo4jBoltDriver.Conn, cm *channelMember) error {
	_, err := conn.ExecNeo(
		`MATCH
			(user:User {userID: {userID}}),
			(channel:Channel {channelID: {channelID}})
		CREATE (user)-[:MEMBER_OF]->(channel)`,
		map[string]interface{}{"channelID": cm.channelID, "userID": cm.userID},
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

func createPostedInChannelVertex(conn golangNeo4jBoltDriver.Conn, p *post) error {
	_, err := conn.ExecNeo("MATCH (post:Post {postID: {postID}}), (channel:Channel {channelID: {channelID}}) CREATE (post)-[:POSTED_IN]->(channel)",
		map[string]interface{}{"postID": p.id, "channelID": p.channelID},
	)
	if err != nil {
		return err
	}
	return nil
}

func createPostedByUserVertex(conn golangNeo4jBoltDriver.Conn, p *post) error {
	_, err := conn.ExecNeo("MATCH (post:Post {postID: {postID}}), (user:User {userID: {userID}}) CREATE (post)-[:POSTED_BY]->(user)",
		map[string]interface{}{"postID": p.id, "userID": p.userID},
	)
	if err != nil {
		return err
	}
	return nil
}

func createChannelPartOfTeamVertex(conn golangNeo4jBoltDriver.Conn, c *channel) error {
	_, err := conn.ExecNeo("MATCH (channel:Channel {channelID: {channelID}}), (team:Team {teamID: {teamID}}) CREATE (channel)-[:PART_OF]->(team)",
		map[string]interface{}{"channelID": c.id, "teamID": c.teamID},
	)
	if err != nil {
		return err
	}
	return nil
}

func createUserCreatedChannelVertex(conn golangNeo4jBoltDriver.Conn, c *channel) error {
	_, err := conn.ExecNeo("MATCH (channel:Channel {channelID: {channelID}}), (user:User {userID: {userID}}) CREATE (user)-[:CREATED]->(channel)",
		map[string]interface{}{"channelID": c.id, "userID": c.creatorID},
	)
	if err != nil {
		return err
	}
	return nil
}

func createReactionVertex(conn golangNeo4jBoltDriver.Conn, r *reaction) error {
	_, err := conn.ExecNeo(
		"MATCH (post:Post {postID: {postID}}), (user:User {userID: {userID}}) CREATE (user)-[:REACTED_TO {emoji:[{emojiName}]}]->(post)",
		map[string]interface{}{"postID": r.postID, "userID": r.userID, "emojiName": r.emojiName},
	)
	if err != nil {
		return err
	}
	return nil
}
