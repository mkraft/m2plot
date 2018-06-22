package cmd

import (
	"database/sql"
)

const batchSize = 100

func teamsIterator(db *sql.DB, batchSize int) func() ([]*team, error) {
	offset := 0
	return func() ([]*team, error) {
		var result []*team
		var err error
		if result, err = teams(db, batchSize, offset); err != nil {
			return nil, err
		}
		offset += batchSize
		return result, nil
	}
}

func forEachTeam(db *sql.DB, itemF func(*team)) error {
	nextTeamsBatch := teamsIterator(db, batchSize)
	var teamBatch []*team
	var err error
	for teamBatch, err = nextTeamsBatch(); len(teamBatch) > 0 && err == nil; teamBatch, err = nextTeamsBatch() {
		for _, team := range teamBatch {
			itemF(team)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func channelsIterator(db *sql.DB, batchSize int) func() ([]*channel, error) {
	offset := 0
	return func() ([]*channel, error) {
		var result []*channel
		var err error
		if result, err = channels(db, batchSize, offset); err != nil {
			return nil, err
		}
		offset += batchSize
		return result, nil
	}
}

func forEachChannel(db *sql.DB, itemF func(*channel)) error {
	nextChannelsBatch := channelsIterator(db, batchSize)
	var channelBatch []*channel
	var err error
	for channelBatch, err = nextChannelsBatch(); len(channelBatch) > 0 && err == nil; channelBatch, err = nextChannelsBatch() {
		for _, channel := range channelBatch {
			itemF(channel)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func usersIterator(db *sql.DB, batchSize int) func() ([]*user, error) {
	offset := 0
	return func() ([]*user, error) {
		var result []*user
		var err error
		if result, err = users(db, batchSize, offset); err != nil {
			return nil, err
		}
		offset += batchSize
		return result, nil
	}
}

func forEachUser(db *sql.DB, itemF func(*user)) error {
	nextUsersBatch := usersIterator(db, batchSize)
	var usersBatch []*user
	var err error
	for usersBatch, err = nextUsersBatch(); len(usersBatch) > 0 && err == nil; usersBatch, err = nextUsersBatch() {
		for _, user := range usersBatch {
			itemF(user)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func postsIterator(db *sql.DB, batchSize int) func() ([]*post, error) {
	offset := 0
	return func() ([]*post, error) {
		var result []*post
		var err error
		if result, err = posts(db, batchSize, offset); err != nil {
			return nil, err
		}
		offset += batchSize
		return result, nil
	}
}

func forEachPost(db *sql.DB, itemF func(*post)) error {
	nextPostsBatch := postsIterator(db, batchSize)
	var postsBatch []*post
	var err error
	for postsBatch, err = nextPostsBatch(); len(postsBatch) > 0 && err == nil; postsBatch, err = nextPostsBatch() {
		for _, post := range postsBatch {
			itemF(post)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
