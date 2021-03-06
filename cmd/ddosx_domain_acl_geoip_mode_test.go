package cmd

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/cli/test/mocks"
	"github.com/ukfast/cli/test/test_output"
	"github.com/ukfast/sdk-go/pkg/service/ddosx"
)

func Test_ddosxDomainACLGeoIPRulesModeShowCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		err := ddosxDomainACLGeoIPRulesModeShowCmd().Args(nil, []string{"testdomain1.co.uk"})

		assert.Nil(t, err)
	})

	t.Run("InvalidArgs_Error", func(t *testing.T) {
		err := ddosxDomainACLGeoIPRulesModeShowCmd().Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing domain", err.Error())
	})
}

func Test_ddosxDomainACLGeoIPRulesModeShow(t *testing.T) {
	t.Run("SingleDomain", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		service.EXPECT().GetDomainACLGeoIPRulesMode("testdomain1.co.uk").Return(ddosx.ACLGeoIPRulesMode(""), nil).Times(1)

		ddosxDomainACLGeoIPRulesModeShow(service, &cobra.Command{}, []string{"testdomain1.co.uk"})
	})

	t.Run("MultipleDomains", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		gomock.InOrder(
			service.EXPECT().GetDomainACLGeoIPRulesMode("testdomain1.co.uk").Return(ddosx.ACLGeoIPRulesMode(""), nil),
			service.EXPECT().GetDomainACLGeoIPRulesMode("testdomain2.co.uk").Return(ddosx.ACLGeoIPRulesMode(""), nil),
		)

		ddosxDomainACLGeoIPRulesModeShow(service, &cobra.Command{}, []string{"testdomain1.co.uk", "testdomain2.co.uk"})
	})

	t.Run("GetDomainACLGeoIPRulesModeError_OutputsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		service.EXPECT().GetDomainACLGeoIPRulesMode("testdomain1.co.uk").Return(ddosx.ACLGeoIPRulesMode(""), errors.New("test error"))

		test_output.AssertErrorOutput(t, "Error retrieving domain [testdomain1.co.uk] ACL GeoIP rules mode: test error\n", func() {
			ddosxDomainACLGeoIPRulesModeShow(service, &cobra.Command{}, []string{"testdomain1.co.uk"})
		})
	})
}

func Test_ddosxDomainACLGeoIPRulesModeUpdateCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		cmd := ddosxDomainACLGeoIPRulesModeUpdateCmd()
		err := cmd.Args(nil, []string{"testdomain1.co.uk"})

		assert.Nil(t, err)
	})

	t.Run("MissingDomain_Error", func(t *testing.T) {
		cmd := ddosxDomainACLGeoIPRulesModeUpdateCmd()
		err := cmd.Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing domain", err.Error())
	})
}

func Test_ddosxDomainACLGeoIPRulesModeUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)
		cmd := ddosxDomainACLGeoIPRulesModeUpdateCmd()
		cmd.Flags().Set("mode", "whitelist")

		expectedRequest := ddosx.PatchACLGeoIPRulesModeRequest{
			Mode: ddosx.ACLGeoIPRulesModeWhitelist,
		}

		gomock.InOrder(
			service.EXPECT().PatchDomainACLGeoIPRulesMode("testdomain1.co.uk", gomock.Eq(expectedRequest)).Return(nil),
			service.EXPECT().GetDomainACLGeoIPRulesMode("testdomain1.co.uk").Return(ddosx.ACLGeoIPRulesMode(""), nil),
		)

		ddosxDomainACLGeoIPRulesModeUpdate(service, cmd, []string{"testdomain1.co.uk"})
	})

	t.Run("InvalidMode_OutputsFatal", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)
		cmd := ddosxDomainACLGeoIPRulesModeUpdateCmd()
		cmd.Flags().Set("mode", "invalidmode")

		test_output.AssertFatalOutputFunc(t, func(stdErr string) {
			assert.Contains(t, stdErr, "Invalid ddosx.ACLGeoIPRulesMode")
		}, func() {
			ddosxDomainACLGeoIPRulesModeUpdate(service, cmd, []string{"testdomain1.co.uk", "00000000-0000-0000-0000-000000000000"})
		})
	})

	t.Run("PatchDomainACLGeoIPRulesMode_OutputsFatal", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		service.EXPECT().PatchDomainACLGeoIPRulesMode("testdomain1.co.uk", gomock.Any()).Return(errors.New("test error"))

		test_output.AssertFatalOutput(t, "Error updating domain ACL GeoIP rule filtering mode: test error\n", func() {
			ddosxDomainACLGeoIPRulesModeUpdate(service, &cobra.Command{}, []string{"testdomain1.co.uk"})
		})
	})

	t.Run("GetDomainACLGeoIPRulesMode_OutputsFatal", func(t *testing.T) {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		gomock.InOrder(
			service.EXPECT().PatchDomainACLGeoIPRulesMode("testdomain1.co.uk", gomock.Any()).Return(nil),
			service.EXPECT().GetDomainACLGeoIPRulesMode("testdomain1.co.uk").Return(ddosx.ACLGeoIPRulesMode(""), errors.New("test error")),
		)

		test_output.AssertFatalOutput(t, "Error retrieving updated domain ACL GeoIP rule filtering mode: test error\n", func() {
			ddosxDomainACLGeoIPRulesModeUpdate(service, &cobra.Command{}, []string{"testdomain1.co.uk"})
		})
	})
}
