package cmd

import (
	"database/sql"
)

const batchSize = 100

func forEachTeam(db *sql.DB, itemF func(*team) error) error {
	offset := 0
	var batch []*team
	var err error
	for batch, err = teams(db, batchSize, offset); len(batch) > 0; batch, err = teams(db, batchSize, offset+batchSize) {
		offset += batchSize
		if err != nil {
			return err
		}
		for _, item := range batch {
			if err = itemF(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func forEachChannel(db *sql.DB, itemF func(*channel) error) error {
	offset := 0
	var batch []*channel
	var err error
	for batch, err = channels(db, batchSize, offset); len(batch) > 0; batch, err = channels(db, batchSize, offset+batchSize) {
		offset += batchSize
		if err != nil {
			return err
		}
		for _, item := range batch {
			if err = itemF(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func forEachUser(db *sql.DB, itemF func(*user) error) error {
	offset := 0
	var batch []*user
	var err error
	for batch, err = users(db, batchSize, offset); len(batch) > 0; batch, err = users(db, batchSize, offset+batchSize) {
		offset += batchSize
		if err != nil {
			return err
		}
		for _, item := range batch {
			if err = itemF(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func forEachPost(db *sql.DB, itemF func(*post) error) error {
	offset := 0
	var batch []*post
	var err error
	for batch, err = posts(db, batchSize, offset); len(batch) > 0; batch, err = posts(db, batchSize, offset+batchSize) {
		offset += batchSize
		if err != nil {
			return err
		}
		for _, item := range batch {
			if err = itemF(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func forEachChannelMember(db *sql.DB, itemF func(*channelMember) error) error {
	offset := 0
	var batch []*channelMember
	var err error
	for batch, err = channelMembers(db, batchSize, offset); len(batch) > 0; batch, err = channelMembers(db, batchSize, offset+batchSize) {
		offset += batchSize
		if err != nil {
			return err
		}
		for _, item := range batch {
			if err = itemF(item); err != nil {
				return err
			}
		}
	}
	return nil
}

func forEachTeamMember(db *sql.DB, itemF func(*teamMember) error) error {
	offset := 0
	var batch []*teamMember
	var err error
	for batch, err = teamMembers(db, batchSize, offset); len(batch) > 0; batch, err = teamMembers(db, batchSize, offset+batchSize) {
		offset += batchSize
		if err != nil {
			return err
		}
		for _, item := range batch {
			if err = itemF(item); err != nil {
				return err
			}
		}
	}
	return nil
}
