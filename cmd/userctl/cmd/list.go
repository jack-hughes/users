package cmd

import (
	"context"
	"github.com/jack-hughes/users/cmd/userctl/utils"
	"github.com/jack-hughes/users/pkg/api/users"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
)

func listCmd() *cobra.Command {
	ctx = context.TODO()
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists users based on a filter of a 2 character country code",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			conn = utils.NewGRPCConn(viper.GetString("host"), viper.GetString("port"))
			client = users.NewUsersClient(conn)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			_ = conn.Close()
		},
		Run: func(cmd *cobra.Command, args []string) {
			req := &users.ListUsersRequest{
				Filter: viper.GetString(utils.Country),
			}

			stream, err := client.List(ctx, req)
			if err != nil {
				log.Fatalf("request failure: %s", err.Error())
			}
			done := make(chan bool)
			go func() {
				for {
					resp, err := stream.Recv()
					if err == io.EOF {
						done <- true
						return
					}
					if err != nil {
						log.Fatal("cannot receive from stream", err)
					}
					utils.ResponsePrinter(resp)
				}
			}()
			<-done
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := cmd.MarkFlagRequired(utils.Country); err != nil {
				log.Fatalf("required flag: %v", utils.Country)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&id, utils.Id, "", "the user id to modify")
	cmd.PersistentFlags().StringVar(&firstName, utils.FirstName, "", "the users first name")
	cmd.PersistentFlags().StringVar(&lastName, utils.LastName, "", "the users last name")
	cmd.PersistentFlags().StringVar(&nickname, utils.Nickname, "", "the users nickname")
	cmd.PersistentFlags().StringVar(&password, utils.Password, "", "the users password")
	cmd.PersistentFlags().StringVar(&email, utils.Email, "", "the users email address")
	cmd.PersistentFlags().StringVar(&country, utils.Country, "", "the users 2 digit country code")
	err := viper.BindPFlags(cmd.PersistentFlags())
	if err != nil {
		log.Fatalf("failed to bind flag values: %v", err)
	}

	return cmd
}
