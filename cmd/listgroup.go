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
	"github.com/spf13/cobra"

	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"context"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"os"
	"time"
)

//var groupId string
// listgroupCmd represents the listgroup command
var listgroupCmd = &cobra.Command{
	Use:   "listgroup",
	Short: "lists a group or groups",
	Long: `list a group or groups`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listgroup called")
		groupId := cmd.Flag("groupid").Value.String()
		//timeout := cmd.Flag("timeout").Value
		fmt.Print(groupId)
		//fmt.Print(timeout)

		sess := session.New()
		creds := credentials.NewChainCredentials(
			//new(credentials.FileProvider),
			new(credentials.EnvProvider),
		)
		svc := elastigroup.New(sess, &spotinst.Config{Credentials: creds})


		var timeout time.Duration


		ctx := context.Background()
		var cancelFn func()
		if timeout > 0 {
			ctx, cancelFn = context.WithTimeout(ctx, timeout)
		}
		// Ensure the context is canceled to prevent leaking.
		// See context package for more information, https://golang.org/pkg/context/
		defer func() {
			if cancelFn != nil {
				cancelFn()
			}
		}()

		// Read group configuration. The Context will interrupt the request if the
		// timeout expires.
		out, err := svc.CloudProviderAWS().Read(ctx, &aws.ReadGroupInput{
			GroupID: spotinst.String(groupId),
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read group, %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("successfully read group %#v\n", out.Group)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listgroupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listgroupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listgroupCmd.Flags().StringP("groupid", "g", "","the group id you wish to check")
	//listgroupCmd.Flags().DurationP("timeout", "d", 0,"how long to wait before timing out")
	rootCmd.AddCommand(listgroupCmd)
}
