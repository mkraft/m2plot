package cmd

import (
	"database/sql"
	"strings"
)

func teams(db *sql.DB, limit, offset int) ([]*team, error) {
	var teams []*team
	rows, err := db.Query("SELECT Id, Name, CreateAt, DeleteAt, Type FROM Teams LIMIT $1 OFFSET $2", limit, offset)
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

func channels(db *sql.DB, limit, offset int) ([]*channel, error) {
	var channels []*channel
	rows, err := db.Query("SELECT Id, Name, CreateAt, DeleteAt, Type, TeamId FROM Channels LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		channel := new(channel)
		if err := rows.Scan(&channel.id, &channel.name, &channel.createAt, &channel.deleteAt, &channel.Type, &channel.teamID); err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	return channels, nil
}

func users(db *sql.DB, limit, offset int) ([]*user, error) {
	var users []*user
	rows, err := db.Query("SELECT Id, Username, CreateAt, DeleteAt FROM Users LIMIT $1 OFFSET $2", limit, offset)
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

func posts(db *sql.DB, limit, offset int) ([]*post, error) {
	var posts []*post
	rows, err := db.Query("SELECT Id, Message, CreateAt, DeleteAt, Hashtags, ChannelId, UserId FROM Posts LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	hashtags := new(string)
	for rows.Next() {
		post := new(post)
		if err := rows.Scan(&post.id, &post.message, &post.createAt, &post.deleteAt, hashtags, &post.channelID, &post.userID); err != nil {
			return nil, err
		}
		post.hashtags = strings.Split(*hashtags, " ")
		posts = append(posts, post)
	}
	return posts, nil
}

func teamMembers(db *sql.DB, limit, offset int) ([]*teamMember, error) {
	var teamMembers []*teamMember
	rows, err := db.Query("SELECT UserId, TeamId, DeleteAt FROM TeamMembers LIMIT $1 OFFSET $2", limit, offset)
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

func channelMembers(db *sql.DB, limit, offset int) ([]*channelMember, error) {
	var channelMembers []*channelMember
	rows, err := db.Query("SELECT UserId, ChannelId FROM ChannelMembers LIMIT $1 OFFSET $2", limit, offset)
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
