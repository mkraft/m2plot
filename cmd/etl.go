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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// etlCmd represents the etl command
var etlCmd = &cobra.Command{
	Use:   "etl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mmAdapter := viper.GetString("mattermost_db.adapter")
		mmConnStr := viper.GetString("mattermost_db.connectionString")
		db, err := sql.Open(mmAdapter, mmConnStr)
		if err != nil {
			return err
		}
		defer db.Close()

		neo4jConnStr := viper.GetString("neo4j.connectionString")
		driver := golangNeo4jBoltDriver.NewDriver()
		conn, err := driver.OpenNeo(neo4jConnStr)
		if err != nil {
			return err
		}
		defer conn.Close()

		forEachTeam(db, func(t *team) {
			err = createTeamEdge(conn, t)
			if err != nil {
				log.Printf("[ERROR] creating edge for team '%v'", t.name)
				log.Print(err)
			}
		})

		forEachChannel(db, func(c *channel) {
			err = createChannelEdge(conn, c)
			if err != nil {
				log.Printf("[ERROR] creating edge for channel '%v'", c.name)
				log.Print(err)
			}
		})

		forEachUser(db, func(u *user) {
			err = createUserEdge(conn, u)
			if err != nil {
				log.Printf("[ERROR] creating edge for user '%v'", u.username)
				log.Print(err)
			}
		})

		forEachPost(db, func(p *post) {
			err = createPostEdge(conn, p)
			if err != nil {
				log.Printf("[ERROR] creating edge for post '%v'", p.id)
				log.Print(err)
			}
		})

		return nil
	},
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
