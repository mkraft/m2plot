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
	nextBatch := teamsIterator(db, batchSize)
	var batch []*team
	var err error
	for batch, err = nextBatch(); len(batch) > 0 && err == nil; batch, err = nextBatch() {
		for _, item := range batch {
			itemF(item)
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
	nextBatch := channelsIterator(db, batchSize)
	var batch []*channel
	var err error
	for batch, err = nextBatch(); len(batch) > 0 && err == nil; batch, err = nextBatch() {
		for _, item := range batch {
			itemF(item)
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
	nextBatch := usersIterator(db, batchSize)
	var batch []*user
	var err error
	for batch, err = nextBatch(); len(batch) > 0 && err == nil; batch, err = nextBatch() {
		for _, item := range batch {
			itemF(item)
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
	nextBatch := postsIterator(db, batchSize)
	var batch []*post
	var err error
	for batch, err = nextBatch(); len(batch) > 0 && err == nil; batch, err = nextBatch() {
		for _, item := range batch {
			itemF(item)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func channelMembersIterator(db *sql.DB, batchSize int) func() ([]*channelMember, error) {
	offset := 0
	return func() ([]*channelMember, error) {
		var result []*channelMember
		var err error
		if result, err = channelMembers(db, batchSize, offset); err != nil {
			return nil, err
		}
		offset += batchSize
		return result, nil
	}
}

func forEachChannelMember(db *sql.DB, itemF func(*channelMember)) error {
	nextBatch := channelMembersIterator(db, batchSize)
	var batch []*channelMember
	var err error
	for batch, err = nextBatch(); len(batch) > 0 && err == nil; batch, err = nextBatch() {
		for _, item := range batch {
			itemF(item)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func teamMembersIterator(db *sql.DB, batchSize int) func() ([]*teamMember, error) {
	offset := 0
	return func() ([]*teamMember, error) {
		var result []*teamMember
		var err error
		if result, err = teamMembers(db, batchSize, offset); err != nil {
			return nil, err
		}
		offset += batchSize
		return result, nil
	}
}

func forEachTeamMember(db *sql.DB, itemF func(*teamMember)) error {
	nextBatch := teamMembersIterator(db, batchSize)
	var batch []*teamMember
	var err error
	for batch, err = nextBatch(); len(batch) > 0 && err == nil; batch, err = nextBatch() {
		for _, item := range batch {
			itemF(item)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
