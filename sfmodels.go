package salesforcehandler

import "fmt"

type SalesForceParam struct {
	Urls          SalesForceUrls
	Autentication SalesForceAutentication
}

type SalesForceUrls struct {
	Autentication string
	Get           string
	Patch         string
}

type SalesForceAutentication struct {
	GrantType    string
	ClientId     string
	ClientSecret string
	UserName     string
	Password     string
}

type SalesForceHandler struct {
	urls        *SalesForceUrls
	accessToken *string
}

func (sf *SalesForceParam) checkdata() error {
	if err := sf.Urls.checkdata(); err != nil {
		return err
	}

	if err := sf.Autentication.checkdata(); err != nil {
		return err
	}

	return nil
}

func (sf *SalesForceUrls) checkdata() error {
	if len(sf.Get) == 0 || len(sf.Patch) == 0 {
		return fmt.Errorf("url's not found")
	}

	return nil
}

func (sf *SalesForceAutentication) checkdata() error {
	if len(sf.GrantType) == 0 || len(sf.ClientId) == 0 || len(sf.ClientSecret) == 0 || len(sf.UserName) == 0 || len(sf.Password) == 0 {
		return fmt.Errorf("salesforce authentication data not found")
	}

	return nil
}

type SalesForcePatchObject struct {
	Name string
	Id   string
	Data any
}

func (sf *SalesForcePatchObject) checkdata() error {
	if len(sf.Name) == 0 || len(sf.Id) == 0 || sf.Data == nil {
		return fmt.Errorf("data for object update not found")
	}

	return nil
}

type SalesForceDownloadFilesParam struct {
	Name string
	Id   string
}

func (sf *SalesForceDownloadFilesParam) checkdata() error {
	if len(sf.Name) == 0 || len(sf.Id) == 0 {
		return fmt.Errorf("data for downloading the file was not found")
	}

	return nil
}
