// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// etlCmd represents the etl command
var etlCmd = &cobra.Command{
	Use:   "etl",
	Short: "Does an initial dump of Mattermost data into Neo4j",
	Long: `Creates the following edges:
* Channels
* Posts
* Teams
* Users
and the following vertices:
* user member of team
* user member of channel
* post posted in channel
* post posted by user
* channel part of team
* user created channel`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sqlConn, err := sql.Open(
			viper.GetString("mattermost_db.adapter"),
			viper.GetString("mattermost_db.connectionString"),
		)
		if err != nil {
			return err
		}
		defer sqlConn.Close()

		graphConn, err := golangNeo4jBoltDriver.NewDriver().OpenNeo(viper.GetString("neo4j.connectionString"))
		if err != nil {
			return err
		}
		defer graphConn.Close()

		if err = forEachTeam(sqlConn, teamF(graphConn)); err != nil {
			return err
		}

		if err = forEachUser(sqlConn, userF(graphConn)); err != nil {
			return err
		}

		if err = forEachChannel(sqlConn, channelF(graphConn)); err != nil {
			return err
		}

		if err = forEachPost(sqlConn, postF(graphConn)); err != nil {
			return err
		}

		if err = forEachTeamMember(sqlConn, teamMemberF(graphConn)); err != nil {
			return err
		}

		if err = forEachChannelMember(sqlConn, channelMemberF(graphConn)); err != nil {
			return err
		}

		return nil
	},
}

func teamF(conn golangNeo4jBoltDriver.Conn) func(*team) error {
	return func(t *team) error {
		err := createTeamEdge(conn, t)
		if err != nil {
			return err
		}
		return nil
	}
}

func userF(conn golangNeo4jBoltDriver.Conn) func(*user) error {
	return func(u *user) error {
		err := createUserEdge(conn, u)
		if err != nil {
			return err
		}
		return nil
	}
}

func channelF(conn golangNeo4jBoltDriver.Conn) func(*channel) error {
	return func(c *channel) error {
		err := createChannelEdge(conn, c)
		if err != nil {
			return err
		}

		err = createChannelPartOfTeamVertex(conn, c)
		if err != nil {
			return err
		}

		err = createUserCreatedChannelVertex(conn, c)
		if err != nil {
			return err
		}

		return nil
	}
}

func postF(conn golangNeo4jBoltDriver.Conn) func(*post) error {
	return func(p *post) error {
		err := createPostEdge(conn, p)
		if err != nil {
			return err
		}

		err = createPostedInChannelVertex(conn, p)
		if err != nil {
			return err
		}

		err = createPostedByUserVertex(conn, p)
		if err != nil {
			return err
		}

		return nil
	}
}

func teamMemberF(conn golangNeo4jBoltDriver.Conn) func(*teamMember) error {
	return func(tm *teamMember) error {
		err := createTeamMemberVertex(conn, tm)
		if err != nil {
			return err
		}
		return nil
	}
}

func channelMemberF(conn golangNeo4jBoltDriver.Conn) func(*channelMember) error {
	return func(cm *channelMember) error {
		err := createChannelMemberVertex(conn, cm)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	rootCmd.AddCommand(etlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// etlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// etlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
