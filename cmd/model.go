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

type team struct {
	id       string
	name     string
	createAt int
	deleteAt int
	Type     string
}

func (t *team) displayType() string {
	var display string
	switch t.Type {
	case "O":
		display = "Open"
	case "I":
		display = "Invite"
	default:
		display = ""
	}
	return display
}

type channel struct {
	id        string
	name      string
	createAt  int
	deleteAt  int
	Type      string
	teamID    string
	creatorID string
}

func (c *channel) displayType() string {
	var display string
	switch c.Type {
	case "O":
		display = "Open"
	case "P":
		display = "Private"
	case "D":
		display = "Direct"
	case "G":
		display = "Group"
	default:
		display = ""
	}
	return display
}

type user struct {
	id       string
	username string
	createAt int
	deleteAt int
}

type post struct {
	id        string
	message   string
	createAt  int
	deleteAt  int
	hashtags  []string
	channelID string
	userID    string
}

type teamMember struct {
	teamID   string
	userID   string
	deleteAt int
}

type channelMember struct {
	channelID string
	userID    string
}
