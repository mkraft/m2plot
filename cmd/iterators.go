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

func forEachPublicChannel(db *sql.DB, itemF func(*channel) error) error {
	offset := 0
	var batch []*channel
	var err error
	for batch, err = publicChannels(db, batchSize, offset); len(batch) > 0; batch, err = publicChannels(db, batchSize, offset+batchSize) {
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

func forEachPublicPost(db *sql.DB, itemF func(*post) error) error {
	offset := 0
	var batch []*post
	var err error
	for batch, err = publicPosts(db, batchSize, offset); len(batch) > 0; batch, err = publicPosts(db, batchSize, offset+batchSize) {
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
	for batch, err = publicChannelMembers(db, batchSize, offset); len(batch) > 0; batch, err = publicChannelMembers(db, batchSize, offset+batchSize) {
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
