package salesforcehandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (sf *SalesForceHandler) Get(query string) (map[string]interface{}, error) {
	if len(query) == 0 {
		return nil, fmt.Errorf("query was not sent.")
	}

	reqUrlAddress := fmt.Sprintf("%s/query?q=%s", sf.urls.Get, query)
	reqAuthorization := fmt.Sprintf("Bearer %s", *sf.accessToken)

	req, err := http.NewRequest(http.MethodGet, reqUrlAddress, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", reqAuthorization)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		bodyErr, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error in request. status code: %d", res.StatusCode)
		}

		return nil, fmt.Errorf("error in request. status code: %d | details: %s", res.StatusCode, string(bodyErr))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (sf *SalesForceHandler) Patch(param *SalesForcePatchObject) error {
	if param == nil {
		return fmt.Errorf("data for object update not found")
	} else if err := param.checkdata(); err != nil {
		return err
	}

	reqUrlAddress := fmt.Sprintf("%s/sobjects/%s/%s", sf.urls.Patch, param.Name, param.Id)
	reqAuthorization := fmt.Sprintf("Bearer %s", *sf.accessToken)

	payload, err := json.Marshal(param.Data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, reqUrlAddress, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", reqAuthorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		bodyErr, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error in request. status code: %d", res.StatusCode)
		}

		return fmt.Errorf("error in request. status code: %d | details: %s", res.StatusCode, string(bodyErr))
	}

	return nil
}

func (sf *SalesForceHandler) DownloadFile(param *SalesForceDownloadFilesParam) ([]byte, error) {
	if param == nil {
		return nil, fmt.Errorf("data for downloading the file was not found")
	} else if err := param.checkdata(); err != nil {
		return nil, err
	}

	reqUrlAddress := fmt.Sprintf("%s/sobjects/%s/Document/%s", sf.urls.Patch, param.Name, param.Id)
	reqAuthorization := fmt.Sprintf("Bearer %s", *sf.accessToken)

	req, err := http.NewRequest(http.MethodGet, reqUrlAddress, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", reqAuthorization)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		bodyErr, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error in request. status code: %d", res.StatusCode)
		}

		return nil, fmt.Errorf("error in request. status code: %d | details: %s", res.StatusCode, string(bodyErr))
	}

	base64Bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler body: %w", err)
	}

	return base64Bytes, nil
}
