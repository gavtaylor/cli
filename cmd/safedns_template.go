package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ukfast/cli/internal/pkg/helper"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/sdk-go/pkg/ptr"
	"github.com/ukfast/sdk-go/pkg/service/safedns"
)

func safednsTemplateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "sub-commands relating to templates",
	}

	// Child commands
	cmd.AddCommand(safednsTemplateListCmd())
	cmd.AddCommand(safednsTemplateShowCmd())
	cmd.AddCommand(safednsTemplateCreateCmd())
	cmd.AddCommand(safednsTemplateUpdateCmd())
	cmd.AddCommand(safednsTemplateDeleteCmd())

	// Child root commands
	cmd.AddCommand(safednsTemplateRecordRootCmd())

	return cmd
}

func safednsTemplateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Lists templates",
		Long:    "This command lists templates",
		Example: "ukfast safedns template list",
		Run: func(cmd *cobra.Command, args []string) {
			safednsTemplateList(getClient().SafeDNSService(), cmd, args)
		},
	}

	cmd.Flags().String("name", "", "Template name for filtering")

	return cmd
}

func safednsTemplateList(service safedns.SafeDNSService, cmd *cobra.Command, args []string) {
	params, err := helper.GetAPIRequestParametersFromFlags(cmd)
	if err != nil {
		output.Fatal(err.Error())
		return
	}

	if cmd.Flags().Changed("name") {
		filterName, _ := cmd.Flags().GetString("name")
		params.WithFilter(helper.GetFilteringInferOperator("name", filterName))
	}

	templates, err := service.GetTemplates(params)
	if err != nil {
		output.Fatalf("Error retrieving templates: %s", err)
		return
	}

	outputSafeDNSTemplates(templates)

}

func safednsTemplateShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "show <template: name/id>...",
		Short:   "Shows a template",
		Long:    "This command shows one or more templates",
		Example: "ukfast safedns template show \"main template\"",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing template")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			safednsTemplateShow(getClient().SafeDNSService(), cmd, args)
		},
	}
}

func safednsTemplateShow(service safedns.SafeDNSService, cmd *cobra.Command, args []string) {
	var templates []safedns.Template
	for _, arg := range args {
		template, err := getSafeDNSTemplateByNameOrID(service, arg)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving template [%s]: %s", arg, err)
			continue
		}

		templates = append(templates, template)
	}

	outputSafeDNSTemplates(templates)
}

func safednsTemplateCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Creates a template",
		Long:    "This command creates a template",
		Example: "ukfast safedns template create --name \"main template\"",
		Run: func(cmd *cobra.Command, args []string) {
			safednsTemplateCreate(getClient().SafeDNSService(), cmd, args)
		},
	}
	cmd.Flags().String("name", "", "Name of template")
	cmd.MarkFlagRequired("name")
	cmd.Flags().Bool("default", false, "Specifies template is default")

	return cmd
}

func safednsTemplateCreate(service safedns.SafeDNSService, cmd *cobra.Command, args []string) {
	templateName, _ := cmd.Flags().GetString("name")
	templateDefault, _ := cmd.Flags().GetBool("default")

	createRequest := safedns.CreateTemplateRequest{
		Name:    templateName,
		Default: templateDefault,
	}

	id, err := service.CreateTemplate(createRequest)
	if err != nil {
		output.Fatalf("Error creating template: %s", err)
		return
	}

	template, err := service.GetTemplate(id)
	if err != nil {
		output.Fatalf("Error retrieving new template: %s", err)
		return
	}

	outputSafeDNSTemplates([]safedns.Template{template})
}

func safednsTemplateUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update <template: name/id>...",
		Short:   "Updates a template",
		Long:    "This command updates one or more templates",
		Example: "ukfast safedns template update \"main template\" --name \"old template\"",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing template")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			safednsTemplateUpdate(getClient().SafeDNSService(), cmd, args)
		},
	}

	cmd.Flags().String("name", "", "Name of template")
	cmd.Flags().Bool("default", false, "Specifies template is default")

	return cmd
}

func safednsTemplateUpdate(service safedns.SafeDNSService, cmd *cobra.Command, args []string) {
	patchRequest := safedns.PatchTemplateRequest{}

	if cmd.Flags().Changed("name") {
		templateName, _ := cmd.Flags().GetString("name")
		patchRequest.Name = templateName
	}
	if cmd.Flags().Changed("default") {
		templateDefault, _ := cmd.Flags().GetBool("default")
		patchRequest.Default = ptr.Bool(templateDefault)
	}

	var templates []safedns.Template

	for _, arg := range args {
		templateID, err := getSafeDNSTemplateIDByNameOrID(service, arg)
		if err != nil {
			output.OutputWithErrorLevel(err.Error())
			continue
		}

		id, err := service.PatchTemplate(templateID, patchRequest)
		if err != nil {
			output.OutputWithErrorLevelf("Error updating template [%s]: %s", arg, err)
			continue
		}

		template, err := service.GetTemplate(id)
		if err != nil {
			output.OutputWithErrorLevelf("Error retrieving updated template: %s", err)
			continue
		}

		templates = append(templates, template)
	}

	outputSafeDNSTemplates(templates)
}

func safednsTemplateDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "delete <template: name/id>...",
		Short:   "Removes a template",
		Long:    "This command removes one or more templates",
		Example: "ukfast safedns template delete \"main template\"",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Missing template")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			safednsTemplateDelete(getClient().SafeDNSService(), cmd, args)
		},
	}
}

func safednsTemplateDelete(service safedns.SafeDNSService, cmd *cobra.Command, args []string) {
	for _, arg := range args {
		templateID, err := getSafeDNSTemplateIDByNameOrID(service, arg)
		if err != nil {
			output.OutputWithErrorLevel(err.Error())
			continue
		}

		err = service.DeleteTemplate(templateID)
		if err != nil {
			output.OutputWithErrorLevelf("Error removing template [%s]: %s", arg, err)
			continue
		}
	}

}
