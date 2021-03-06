package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/account"
)

func accountContactRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contact",
		Short: "sub-commands relating to contacts",
	}

	// Child commands
	cmd.AddCommand(accountContactListCmd())
	cmd.AddCommand(accountContactShowCmd())

	return cmd
}

func accountContactListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Lists contacts",
		Long:    "This command lists contacts",
		Example: "ukfast account contact list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return accountContactList(getClient().AccountService(), cmd, args)
		},
	}
}

func accountContactList(service account.AccountService, cmd *cobra.Command, args []string) error {
	params, err := helper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		return err
	}

	contacts, err := service.GetContacts(params)
	if err != nil {
		return fmt.Errorf("Error retrieving contacts: %s", err)
	}

	outputAccountContacts(contacts)
	return nil
}

func accountContactShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "show <contact: id>...",
		Short:   "Shows a contact",
		Long:    "This command shows one or more contacts",
		Example: "ukfast account contact show 123",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing contact")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			accountContactShow(getClient().AccountService(), cmd, args)
		},
	}
}

func accountContactShow(service account.AccountService, cmd *cobra.Command, args []string) {
	var contacts []account.Contact
	for _, arg := range args {
		contactID, err := strconv.Atoi(arg)
		if err != nil {
			output.OutputWithErrorLevelf("Invalid contact ID [%s]", arg)
			continue
		}

		contact, err := service.GetContact(contactID)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving contact [%s]: %s", arg, err)
			continue
		}

		contacts = append(contacts, contact)
	}

	outputAccountContacts(contacts)
}
