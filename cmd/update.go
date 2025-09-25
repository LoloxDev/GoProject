package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var (
		id       uint64
		name     string
		email    string
		password string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Met à jour un contact existant",
		Run: func(cmd *cobra.Command, args []string) {
			if id == 0 {
				exitWithError(cmd, fmt.Errorf("un identifiant valide est requis"))
			}

			service, err := getService()
			if err != nil {
				exitWithError(cmd, err)
			}

			contact, err := service.Update(uint(id), name, email, password)
			if err != nil {
				exitWithError(cmd, err)
			}

			cmd.Printf("Contact #%d mis à jour: %s <%s>\n", contact.ID, contact.Name, contact.Email)
		},
	}

	cmd.Flags().Uint64VarP(&id, "id", "i", 0, "Identifiant du contact")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Nouveau nom")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Nouvel email")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Nouveau mot de passe")

	return cmd
}
