package services

import (
	"github.com/h4rimu/kaspro-sdk/logging"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var log = logging.MustGetLogger("kaspro-sdk")

type ThirdParty interface {
	SetApiUrl() (*url.URL, *error)
	HitClient(url url.URL) (*ClientResponse, *error)
}

type ClientParty struct {
	UniqueID    string
	ClientName  string
	HttpMethod  string
	UrlApi      string
	HttpClient  http.Client
	Headers     []map[string]string
	RequestBody io.Reader
}

type ClientResponse struct {
	HttpCode     int
	ByteResponse []byte
}

func (c *ClientParty) SetApiUrl() (*url.URL, *error) {
	url, err := url.Parse(c.UrlApi)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}
	return url, nil
}

func (c *ClientParty) HitClient(url url.URL) (*ClientResponse, *error) {

	request, err := http.NewRequest(c.HttpMethod, url.String(), c.RequestBody)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	for _, element := range c.Headers {
		for key, value := range element {
			if strings.ToLower(key) == "baseauth" {
				baseAuth := strings.Split(value, ":")
				request.SetBasicAuth(baseAuth[0], baseAuth[1])
			} else {
				request.Header.Set(key, value)
			}
		}
	}

	log.Infof(c.UniqueID, "===================================================== HIT "+strings.ToUpper(c.ClientName)+" API =======================================================")
	log.Infof(c.UniqueID, "HTTP METHOD %s ", request.Method)
	log.Infof(c.UniqueID, "URL %s ", request.URL)
	log.Infof(c.UniqueID, "HEADER %s ", request.Header)
	if request.Body != nil {
		log.Infof(c.UniqueID, "BODY %s ", request.Body)
	}
	log.Infof(c.UniqueID, "==========================================================================================================================")

	response, err := c.HttpClient.Do(request)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	byteResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	log.Infof(c.UniqueID, "================================================ RESPONSE "+strings.ToUpper(c.ClientName)+" API ================================================")
	log.Infof(c.UniqueID, "GET RESPONSE FROM "+strings.ToUpper(c.ClientName)+" API %s", string(byteResult))
	log.Infof(c.UniqueID, "======================================================================================================================")

	clientResponse := ClientResponse{
		HttpCode:     response.StatusCode,
		ByteResponse: byteResult,
	}
	return &clientResponse, nil
}
