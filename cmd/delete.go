package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var id uint64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Supprime un contact",
		Run: func(cmd *cobra.Command, args []string) {
			if id == 0 {
				exitWithError(cmd, fmt.Errorf("un identifiant valide est requis"))
			}

			service, err := getService()
			if err != nil {
				exitWithError(cmd, err)
			}

			if err := service.Delete(uint(id)); err != nil {
				exitWithError(cmd, err)
			}

			cmd.Printf("Contact #%d supprimé\n", id)
		},
	}

	cmd.Flags().Uint64VarP(&id, "id", "i", 0, "Identifiant du contact")

	return cmd
}
