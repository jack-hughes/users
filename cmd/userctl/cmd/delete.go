package cmd

import (
	"context"
	"github.com/jack-hughes/users/cmd/userctl/utils"
	"github.com/jack-hughes/users/pkg/api/users"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func deleteCmd() *cobra.Command {
	ctx = context.TODO()
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a user based on their id",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			conn = utils.NewGRPCConn(viper.GetString("host"), viper.GetString("port"))
			client = users.NewUsersClient(conn)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			_ = conn.Close()
		},
		Run: func(cmd *cobra.Command, args []string) {
			req := &users.User{
				Id: viper.GetString(utils.Id),
			}

			_, err := client.Delete(ctx, req)
			if err != nil {
				log.Fatalf("request failure: %s", err.Error())
			}

			log.Printf("deleted user: %s", req.Id)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := cmd.MarkFlagRequired(utils.Id); err != nil {
				log.Fatalf("required flag: %v", utils.Id)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&id, utils.Id, "", "the user id to delete")
	err := viper.BindPFlags(cmd.PersistentFlags())
	if err != nil {
		log.Fatalf("failed to bind flag values: %v", err)
	}

	return cmd
}
