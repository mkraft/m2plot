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

import (
	"database/sql"
	"fmt"
	"strings"
)

func teams(db *sql.DB, limit, offset int) ([]*team, error) {
	var teams []*team
	query := fmt.Sprintf(`
		SELECT Id, Name, CreateAt, DeleteAt, Type 
		FROM Teams 
		ORDER BY CreateAt DESC
		LIMIT %v 
		OFFSET %v 
		`, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		team := new(team)
		if err := rows.Scan(&team.id, &team.name, &team.createAt, &team.deleteAt, &team.Type); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func publicChannels(db *sql.DB, limit, offset int) ([]*channel, error) {
	var channels []*channel
	query := fmt.Sprintf(`
		SELECT Id, Name, CreateAt, DeleteAt, Type, TeamId, CreatorId 
		FROM Channels 
		WHERE Type = 'O' 
		ORDER BY CreateAt DESC
		LIMIT %v 
		OFFSET %v
		`, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		channel := new(channel)
		if err := rows.Scan(&channel.id, &channel.name, &channel.createAt, &channel.deleteAt, &channel.Type,
			&channel.teamID, &channel.creatorID); err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	return channels, nil
}

func users(db *sql.DB, limit, offset int) ([]*user, error) {
	var users []*user
	query := fmt.Sprintf(`
		SELECT Id, Username, CreateAt, DeleteAt 
		FROM Users 
		ORDER BY CreateAt DESC
		LIMIT %v 
		OFFSET %v
		`, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := new(user)
		if err := rows.Scan(&user.id, &user.username, &user.createAt, &user.deleteAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func publicPosts(db *sql.DB, limit, offset int) ([]*post, error) {
	var posts []*post
	query := fmt.Sprintf(`
		SELECT Posts.Id, Message, Posts.CreateAt, Posts.DeleteAt, Hashtags, ChannelId, UserId 
		FROM Posts 
		JOIN Channels ON Channels.Id = Posts.ChannelId 
		WHERE Channels.Type = 'O' AND Posts.Type = '' 
		ORDER BY Posts.CreateAt DESC
		LIMIT %v 
		OFFSET %v
		`, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	hashtags := new(string)
	for rows.Next() {
		post := new(post)
		if err := rows.Scan(&post.id, &post.message, &post.createAt, &post.deleteAt, hashtags, &post.channelID,
			&post.userID); err != nil {
			return nil, err
		}
		post.hashtags = strings.Split(*hashtags, " ")
		posts = append(posts, post)
	}
	return posts, nil
}

func teamMembers(db *sql.DB, limit, offset int) ([]*teamMember, error) {
	var teamMembers []*teamMember
	query := fmt.Sprintf(`
		SELECT UserId, TeamId, DeleteAt 
		FROM TeamMembers 
		ORDER BY TeamId DESC
		LIMIT %v 
		OFFSET %v
		`, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		teamMember := new(teamMember)
		if err := rows.Scan(&teamMember.userID, &teamMember.teamID, &teamMember.deleteAt); err != nil {
			return nil, err
		}
		teamMembers = append(teamMembers, teamMember)
	}
	return teamMembers, nil
}

func publicChannelMembers(db *sql.DB, limit, offset int) ([]*channelMember, error) {
	var channelMembers []*channelMember
	query := fmt.Sprintf(`
		SELECT UserId, ChannelId 
		FROM ChannelMembers 
		JOIN Channels ON Channels.Id = ChannelMembers.ChannelId 
		WHERE Channels.Type = 'O' 
		ORDER BY ChannelMembers.LastUpdateAt DESC
		LIMIT %v 
		OFFSET %v
		`, limit, offset)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		channelMember := new(channelMember)
		if err := rows.Scan(&channelMember.userID, &channelMember.channelID); err != nil {
			return nil, err
		}
		channelMembers = append(channelMembers, channelMember)
	}
	return channelMembers, nil
}
