package ddosx

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSSLs retrieves a list of ssls
func (s *Service) GetSSLs(parameters connection.APIRequestParameters) ([]SSL, error) {
	r := connection.RequestAll{}

	var ssls []SSL
	r.GetNext = func(parameters connection.APIRequestParameters) (connection.ResponseBody, error) {
		response, err := s.getSSLsPaginatedResponseBody(parameters)
		if err != nil {
			return nil, err
		}

		for _, ssl := range response.Data {
			ssls = append(ssls, ssl)
		}

		return response, nil
	}

	err := r.Invoke(parameters)

	return ssls, err
}

// GetSSLsPaginated retrieves a paginated list of ssls
func (s *Service) GetSSLsPaginated(parameters connection.APIRequestParameters) ([]SSL, error) {
	body, err := s.getSSLsPaginatedResponseBody(parameters)

	return body.Data, err
}

func (s *Service) getSSLsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSSLsResponseBody, error) {
	body := &GetSSLsResponseBody{}

	response, err := s.connection.Get("/ddosx/v1/ssls", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse([]int{200}, body)
}

// GetSSL retrieves a single ssl by id
func (s *Service) GetSSL(sslID string) (SSL, error) {
	body, err := s.getSSLResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getSSLResponseBody(sslID string) (*GetSSLResponseBody, error) {
	body := &GetSSLResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	if response.StatusCode == 404 {
		return body, &SSLNotFoundError{ID: sslID}
	}

	return body, response.HandleResponse([]int{200}, body)
}

// CreateSSL retrieves creates an SSL
func (s *Service) CreateSSL(req CreateSSLRequest) (string, error) {
	body, err := s.createSSLResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createSSLResponseBody(req CreateSSLRequest) (*GetSSLResponseBody, error) {
	body := &GetSSLResponseBody{}

	response, err := s.connection.Post("/ddosx/v1/ssls", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse([]int{201}, body)
}

// PatchSSL retrieves patches an SSL
func (s *Service) PatchSSL(sslID string, req PatchSSLRequest) (string, error) {
	body, err := s.patchSSLResponseBody(sslID, req)

	return body.Data.ID, err
}

func (s *Service) patchSSLResponseBody(sslID string, req PatchSSLRequest) (*GetSSLResponseBody, error) {
	body := &GetSSLResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), &req)
	if err != nil {
		return body, err
	}

	if response.StatusCode == 404 {
		return body, &SSLNotFoundError{ID: sslID}
	}

	return body, response.HandleResponse([]int{200}, body)
}

// DeleteSSL deletes patches an SSL
func (s *Service) DeleteSSL(sslID string) error {
	_, err := s.deleteSSLResponseBody(sslID)

	return err
}

func (s *Service) deleteSSLResponseBody(sslID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), nil)
	if err != nil {
		return body, err
	}

	if response.StatusCode == 404 {
		return body, &SSLNotFoundError{ID: sslID}
	}

	return body, response.HandleResponse([]int{204}, body)
}

// GetSSLContent retrieves a single ssl by id
func (s *Service) GetSSLContent(sslID string) (SSLContent, error) {
	body, err := s.getSSLContentResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getSSLContentResponseBody(sslID string) (*GetSSLContentResponseBody, error) {
	body := &GetSSLContentResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s/certificates", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	if response.StatusCode == 404 {
		return body, &SSLNotFoundError{ID: sslID}
	}

	return body, response.HandleResponse([]int{200}, body)
}

// GetSSLPrivateKey retrieves a single ssl by id
func (s *Service) GetSSLPrivateKey(sslID string) (SSLPrivateKey, error) {
	body, err := s.getSSLPrivateKeyResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getSSLPrivateKeyResponseBody(sslID string) (*GetSSLPrivateKeyResponseBody, error) {
	body := &GetSSLPrivateKeyResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s/private-key", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	if response.StatusCode == 404 {
		return body, &SSLNotFoundError{ID: sslID}
	}

	return body, response.HandleResponse([]int{200}, body)
}