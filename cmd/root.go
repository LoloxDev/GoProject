package cmd

import (
	"fmt"
	"os"

	"GoProject/internal/config"
	"GoProject/internal/contacts"
	"GoProject/internal/storage"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Mini-CRM CLI",
}

var service *contacts.Service

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Erreur config:", err)
		os.Exit(1)
	}

	store, err := storage.NewFromConfig(cfg)
	if err != nil {
		fmt.Println("Erreur storage:", err)
		os.Exit(1)
	}
	if err := store.Init(); err != nil {
		fmt.Println("Erreur init storage:", err)
		os.Exit(1)
	}

	service = contacts.NewService(store)

	rootCmd.AddCommand(newAddCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newUpdateCmd())
	rootCmd.AddCommand(newDeleteCmd())
}

func getService() (*contacts.Service, error) {
	if service == nil {
		return nil, fmt.Errorf("service non initialisé")
	}
	return service, nil
}

func exitWithError(cmd *cobra.Command, err error) {
	fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
	os.Exit(1)
}
