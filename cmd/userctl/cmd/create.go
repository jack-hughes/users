package cmd

import (
	"context"
	"github.com/jack-hughes/users/cmd/userctl/utils"
	"github.com/jack-hughes/users/pkg/api/users"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

var (
	conn   *grpc.ClientConn
	ctx    context.Context
	client users.UsersClient

	id        string
	firstName string
	lastName  string
	nickname  string
	password  string
	email     string
	country   string
)

func createCmd() *cobra.Command {
	ctx = context.TODO()
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a user",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			conn = utils.NewGRPCConn(viper.GetString("host"), viper.GetString("port"))
			client = users.NewUsersClient(conn)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			_ = conn.Close()
		},
		Run: func(cmd *cobra.Command, args []string) {
			req := &users.User{
				FirstName: viper.GetString(utils.FirstName),
				LastName:  viper.GetString(utils.LastName),
				Nickname:  viper.GetString(utils.Nickname),
				Password:  viper.GetString(utils.Password),
				Email:     viper.GetString(utils.Email),
				Country:   viper.GetString(utils.Country),
			}

			usr, err := client.Create(ctx, req)
			if err != nil {
				log.Fatalf("request failure: %s", err.Error())
			}

			utils.ResponsePrinter(usr)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := cmd.MarkFlagRequired(utils.FirstName); err != nil {
				log.Fatalf("required flag: %v", utils.FirstName)
			}
			if err := cmd.MarkFlagRequired(utils.LastName); err != nil {
				log.Fatalf("required flag: %v", utils.LastName)
			}
			if err := cmd.MarkFlagRequired(utils.Nickname); err != nil {
				log.Fatalf("required flag: %v", utils.Nickname)
			}
			if err := cmd.MarkFlagRequired(utils.Password); err != nil {
				log.Fatalf("required flag: %v", utils.Password)
			}
			if err := cmd.MarkFlagRequired(utils.Email); err != nil {
				log.Fatalf("required flag: %v", utils.Email)
			}
			if err := cmd.MarkFlagRequired(utils.Country); err != nil {
				log.Fatalf("required flag: %v", utils.Country)
			}
		},
	}

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
