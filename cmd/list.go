package cmd

import (
	"sort"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Affiche tous les contacts enregistrés",
		Run: func(cmd *cobra.Command, args []string) {
			service, err := getService()
			if err != nil {
				exitWithError(cmd, err)
			}

			contacts, err := service.List()
			if err != nil {
				exitWithError(cmd, err)
			}

			if len(contacts) == 0 {
				cmd.Println("Aucun contact enregistré pour le moment.")
				return
			}

			sort.Slice(contacts, func(i, j int) bool {
				return contacts[i].ID < contacts[j].ID
			})

			for _, contact := range contacts {
				cmd.Printf("#%d - %s <%s>\n", contact.ID, contact.Name, contact.Email)
			}
		},
	}

	return cmd
}
