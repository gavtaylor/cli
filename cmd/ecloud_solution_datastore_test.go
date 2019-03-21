package cmd

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/cli/internal/pkg/output"
	"github.com/ukfast/cli/test"
	"github.com/ukfast/cli/test/mocks"
	"github.com/ukfast/sdk-go/pkg/service/ecloud"
)

func Test_ecloudSolutionDatastoreListCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		err := ecloudSolutionDatastoreListCmd().Args(nil, []string{"123"})

		assert.Nil(t, err)
	})

	t.Run("InvalidArgs_Error", func(t *testing.T) {
		err := ecloudSolutionDatastoreListCmd().Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing solution", err.Error())
	})
}

func Test_ecloudSolutionDatastoreList(t *testing.T) {
	t.Run("DefaultRetrieve", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetSolutionDatastores(123, gomock.Any()).Return([]ecloud.Datastore{}, nil).Times(1)

		ecloudSolutionDatastoreList(service, &cobra.Command{}, []string{"123"})
	})

	t.Run("MalformedFlag_OutputsFatal", func(t *testing.T) {
		code := 0
		oldOutputExit := output.SetOutputExit(func(c int) {
			code = c
		})
		defer func() { output.SetOutputExit(oldOutputExit) }()
		defer func() { flagFilter = nil }()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)
		flagFilter = []string{"invalidfilter"}

		output := test.CatchStdErr(t, func() {
			ecloudSolutionDatastoreList(service, &cobra.Command{}, []string{"123"})
		})

		assert.Equal(t, 1, code)
		assert.Equal(t, "Missing value for filtering\n", output)
	})

	t.Run("InvalidSolutionID_OutputsFatal", func(t *testing.T) {
		code := 0
		oldOutputExit := output.SetOutputExit(func(c int) {
			code = c
		})
		defer func() { output.SetOutputExit(oldOutputExit) }()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		output := test.CatchStdErr(t, func() {
			ecloudSolutionDatastoreList(service, &cobra.Command{}, []string{"abc"})
		})

		assert.Equal(t, "Invalid solution ID [abc]\n", output)
		assert.Equal(t, 1, code)
	})

	t.Run("GetDatastoresError_OutputsFatal", func(t *testing.T) {
		code := 0
		oldOutputExit := output.SetOutputExit(func(c int) {
			code = c
		})
		defer func() { output.SetOutputExit(oldOutputExit) }()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockECloudService(mockCtrl)

		service.EXPECT().GetSolutionDatastores(123, gomock.Any()).Return([]ecloud.Datastore{}, errors.New("test error 1")).Times(1)

		output := test.CatchStdErr(t, func() {
			ecloudSolutionDatastoreList(service, &cobra.Command{}, []string{"123"})
		})

		assert.Equal(t, 1, code)
		assert.Equal(t, "Error retrieving solution datastores: test error 1\n", output)
	})
}