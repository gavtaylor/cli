package cmd

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/cli/test/mocks"
	"github.com/ukfast/cli/test/test_output"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func Test_ecloudApplianceList(t *testing.T) {
	t.Run("DefaultRetrieve", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetAppliances(gomock.Any()).Return([]ecloud.Appliance{}, nil).Times(1)

		ecloudApplianceList(service, &cobra.Command{}, []string{})
	})

	t.Run("MalformedFlag_OutputsFatal", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)
		cmd := &cobra.Command{}
		cmd.Flags().StringArray("filter", []string{"invalidfilter"}, "")

		test_output.AssertFatalOutput(t, "Missing value for filtering\n", func() {
			ecloudApplianceList(service, cmd, []string{})
		})
	})

	t.Run("GetAppliancesError_OutputsFatal", func(t *testing.T) {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetAppliances(gomock.Any()).Return([]ecloud.Appliance{}, errors.New("test error")).Times(1)

		test_output.AssertFatalOutput(t, "Error retrieving appliances: test error\n", func() {
			ecloudApplianceList(service, &cobra.Command{}, []string{})
		})
	})
}

func Test_ecloudApplianceShowCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		err := ecloudApplianceShowCmd().Args(nil, []string{"00000000-0000-0000-0000-000000000000"})

		assert.Nil(t, err)
	})

	t.Run("InvalidArgs_Error", func(t *testing.T) {
		err := ecloudApplianceShowCmd().Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing appliance", err.Error())
	})
}

func Test_ecloudApplianceShow(t *testing.T) {
	t.Run("SingleAppliance", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetAppliance("00000000-0000-0000-0000-000000000000").Return(ecloud.Appliance{}, nil).Times(1)

		ecloudApplianceShow(service, &cobra.Command{}, []string{"00000000-0000-0000-0000-000000000000"})
	})

	t.Run("MultipleAppliances", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		gomock.InOrder(
			service.EXPECT().GetAppliance("00000000-0000-0000-0000-000000000000").Return(ecloud.Appliance{}, nil),
			service.EXPECT().GetAppliance("00000000-0000-0000-0000-000000000001").Return(ecloud.Appliance{}, nil),
		)

		ecloudApplianceShow(service, &cobra.Command{}, []string{"00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001"})
	})

	t.Run("GetApplianceError_OutputsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetAppliance("00000000-0000-0000-0000-000000000000").Return(ecloud.Appliance{}, errors.New("test error"))

		test_output.AssertErrorOutput(t, "Error retrieving appliance [00000000-0000-0000-0000-000000000000]: test error\n", func() {
			ecloudApplianceShow(service, &cobra.Command{}, []string{"00000000-0000-0000-0000-000000000000"})
		})
	})
}
