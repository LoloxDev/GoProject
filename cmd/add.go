package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	var name string
	var email string
	var password string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Ajoute un nouveau contact",
		Run: func(cmd *cobra.Command, args []string) {
			service, err := getService()
			if err != nil {
				exitWithError(cmd, err)
			}

			if name == "" || email == "" {
				exitWithError(cmd, fmt.Errorf("les champs name et email sont obligatoires"))
			}

			contact, err := service.Add(name, email, password)
			if err != nil {
				exitWithError(cmd, err)
			}

			cmd.Printf("Contact #%d ajouté: %s <%s>\n", contact.ID, contact.Name, contact.Email)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Nom du contact")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email du contact")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Mot de passe du contact")

	return cmd
}
