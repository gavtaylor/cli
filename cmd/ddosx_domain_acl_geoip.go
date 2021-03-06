package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/service/ddosx"
)

func ddosxDomainACLGeoIPRuleRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "geoip",
		Short: "sub-commands relating to domain ACL GeoIP rules",
	}

	// Child commands
	cmd.AddCommand(ddosxDomainACLGeoIPRuleListCmd())
	cmd.AddCommand(ddosxDomainACLGeoIPRuleShowCmd())
	cmd.AddCommand(ddosxDomainACLGeoIPRuleCreateCmd())
	cmd.AddCommand(ddosxDomainACLGeoIPRuleUpdateCmd())
	cmd.AddCommand(ddosxDomainACLGeoIPRuleDeleteCmd())

	// Child root commands
	cmd.AddCommand(ddosxDomainACLGeoIPRulesModeRootCmd())

	return cmd
}

func ddosxDomainACLGeoIPRuleListCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list <domain: name>",
		Short:   "Lists ACL GeoIP rules",
		Long:    "This command lists domain ACL GeoIP rules",
		Example: "ukfast ddosx domain acl geoip list",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing domain")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ddosxDomainACLGeoIPRuleList(getClient().DDoSXService(), cmd, args)
		},
	}
}

func ddosxDomainACLGeoIPRuleList(service ddosx.DDoSXService, cmd *cobra.Command, args []string) {
	params, err := helper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		output.Fatal(err.Error())
		return
	}

	domains, err := service.GetDomainACLGeoIPRules(args[0], params)
	if err != nil {
		output.Fatalf("Error retrieving domain ACL GeoIP rules: %s", err)
		return
	}

	outputDDoSXACLGeoIPRules(domains)
}

func ddosxDomainACLGeoIPRuleShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "show <domain: name> <rule: id>...",
		Short:   "Shows domain ACL GeoIP rules",
		Long:    "This command shows an ACL GeoIP rule",
		Example: "ukfast ddosx domain acl geoip show example.com 00000000-0000-0000-0000-000000000000",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing domain")
			}
			if len(args) < 2 {
				return errors.New("Missing rule")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ddosxDomainACLGeoIPRuleShow(getClient().DDoSXService(), cmd, args)
		},
	}
}

func ddosxDomainACLGeoIPRuleShow(service ddosx.DDoSXService, cmd *cobra.Command, args []string) {

	var rules []ddosx.ACLGeoIPRule

	for _, arg := range args[1:] {
		rule, err := service.GetDomainACLGeoIPRule(args[0], arg)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving domain ACL GeoIP rule [%s]: %s", arg, err.Error())
			continue
		}

		rules = append(rules, rule)
	}

	outputDDoSXACLGeoIPRules(rules)
}

func ddosxDomainACLGeoIPRuleCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create <domain: name>",
		Short:   "Creates ACL GeoIP rules",
		Long:    "This command creates domain ACL GeoIP rules",
		Example: "ukfast ddosx domain acl geoip create",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing domain")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ddosxDomainACLGeoIPRuleCreate(getClient().DDoSXService(), cmd, args)
		},
	}

	cmd.Flags().String("code", "", "Country code for GeoIP ACL rule")
	cmd.MarkFlagRequired("code")

	return cmd
}

func ddosxDomainACLGeoIPRuleCreate(service ddosx.DDoSXService, cmd *cobra.Command, args []string) {

	createRequest := ddosx.CreateACLGeoIPRuleRequest{}
	createRequest.Code, _ = cmd.Flags().GetString("code")

	id, err := service.CreateDomainACLGeoIPRule(args[0], createRequest)
	if err != nil {
		output.Fatalf("Error creating domain ACL GeoIP rule: %s", err)
		return
	}

	rule, err := service.GetDomainACLGeoIPRule(args[0], id)
	if err != nil {
		output.Fatalf("Error retrieving new domain ACL GeoIP rule [%s]: %s", id, err)
		return
	}

	outputDDoSXACLGeoIPRules([]ddosx.ACLGeoIPRule{rule})
}

func ddosxDomainACLGeoIPRuleUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update <domain: name> <rule: id>...",
		Short:   "Updates ACL GeoIP rules",
		Long:    "This command updates one or more domain ACL GeoIP rules",
		Example: "ukfast ddosx domain acl geoip update example.com 00000000-0000-0000-0000-000000000000 --code GB",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing domain")
			}
			if len(args) < 2 {
				return errors.New("Missing rule")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ddosxDomainACLGeoIPRuleUpdate(getClient().DDoSXService(), cmd, args)
		},
	}

	cmd.Flags().String("code", "", "Country code for GeoIP ACL rule (ISO_3166-1_alpha-2 format)")

	return cmd
}

func ddosxDomainACLGeoIPRuleUpdate(service ddosx.DDoSXService, cmd *cobra.Command, args []string) {
	patchRequest := ddosx.PatchACLGeoIPRuleRequest{}

	if cmd.Flags().Changed("code") {
		patchRequest.Code, _ = cmd.Flags().GetString("code")
	}

	var rules []ddosx.ACLGeoIPRule
	for _, arg := range args[1:] {
		err := service.PatchDomainACLGeoIPRule(args[0], arg, patchRequest)
		if err != nil {
			output.OutputWithErrorLevelf("Error updating domain ACL GeoIP rule [%s]: %s", arg, err.Error())
			continue
		}

		rule, err := service.GetDomainACLGeoIPRule(args[0], arg)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving updated domain ACL GeoIP rule [%s]: %s", arg, err)
			continue
		}

		rules = append(rules, rule)
	}

	outputDDoSXACLGeoIPRules(rules)
}

func ddosxDomainACLGeoIPRuleDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "delete <domain: name> <rule: id>...",
		Short:   "Deletes ACL GeoIP rules",
		Long:    "This command deletes one or more domain ACL GeoIP rules",
		Example: "ukfast ddosx domain acl geoip delete example.com 00000000-0000-0000-0000-000000000000",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing domain")
			}
			if len(args) < 2 {
				return errors.New("Missing rule")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ddosxDomainACLGeoIPRuleDelete(getClient().DDoSXService(), cmd, args)
		},
	}
}

func ddosxDomainACLGeoIPRuleDelete(service ddosx.DDoSXService, cmd *cobra.Command, args []string) {
	for _, arg := range args[1:] {
		err := service.DeleteDomainACLGeoIPRule(args[0], arg)
		if err != nil {
			output.OutputWithErrorLevelf("Error removing domain ACL GeoIP rule [%s]: %s", arg, err.Error())
			continue
		}
	}
}
