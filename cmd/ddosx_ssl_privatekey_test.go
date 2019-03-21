package cmd

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/ukfast/cli/test"
	"github.com/ukfast/cli/test/mocks"
	"github.com/ukfast/sdk-go/pkg/service/ddosx"
)

func Test_ddosxSSLPrivateKeyCmd_Args(t *testing.T) {
	t.Run("ValidArgs_NoError", func(t *testing.T) {
		err := ddosxSSLPrivateKeyShowCmd().Args(nil, []string{"00000000-0000-0000-0000-000000000000"})

		assert.Nil(t, err)
	})

	t.Run("InvalidArgs_Error", func(t *testing.T) {
		err := ddosxSSLPrivateKeyShowCmd().Args(nil, []string{})

		assert.NotNil(t, err)
		assert.Equal(t, "Missing ssl", err.Error())
	})
}

func Test_ddosxSSLPrivateKey(t *testing.T) {
	t.Run("SingleSSL", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		service.EXPECT().GetSSLPrivateKey("00000000-0000-0000-0000-000000000000").Return(ddosx.SSLPrivateKey{}, nil).Times(1)

		ddosxSSLPrivateKeyShow(service, &cobra.Command{}, []string{"00000000-0000-0000-0000-000000000000"})
	})

	t.Run("MultipleSSLs", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		gomock.InOrder(
			service.EXPECT().GetSSLPrivateKey("00000000-0000-0000-0000-000000000000").Return(ddosx.SSLPrivateKey{}, nil),
			service.EXPECT().GetSSLPrivateKey("00000000-0000-0000-0000-000000000001").Return(ddosx.SSLPrivateKey{}, nil),
		)

		ddosxSSLPrivateKeyShow(service, &cobra.Command{}, []string{"00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000001"})
	})

	t.Run("GetSSLPrivateKeyError_OutputsError", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := mocks.NewMockDDoSXService(mockCtrl)

		service.EXPECT().GetSSLPrivateKey("00000000-0000-0000-0000-000000000000").Return(ddosx.SSLPrivateKey{}, errors.New("test error"))

		output := test.CatchStdErr(t, func() {
			ddosxSSLPrivateKeyShow(service, &cobra.Command{}, []string{"00000000-0000-0000-0000-000000000000"})
		})

		assert.Equal(t, "Error retrieving ssl [00000000-0000-0000-0000-000000000000]: test error\n", output)
	})
}
