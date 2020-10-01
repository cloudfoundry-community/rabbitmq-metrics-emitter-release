package management

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/starkandwayne/rabbitmq-metrics-emitter/config"
	"gopkg.in/yaml.v3"
)

type QueueInfo struct {
	Name          string  `yaml:"name"`
	MessagesReady float64 `yaml:"messages_ready"`
}

type VhostInfo struct {
	Name string `yaml:"name"`
}

type Client struct {
	config     *config.Config
	logger     lager.Logger
	httpClient *http.Client
}

func NewManagementClient(logger lager.Logger, config *config.Config) (Client, error) {
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.RMQManagement.SkipSslValidation,
			},
		},
		Timeout: time.Second * 30,
	}
	return Client{
		logger:     logger.Session("management"),
		httpClient: &httpClient,
		config:     config,
	}, nil
}

func (client *Client) GetVhosts() ([]VhostInfo, error) {
	client.logger.Info("get-vhosts")
	var infos []VhostInfo

	req, err := http.NewRequest(http.MethodGet, client.config.RMQManagement.Endpoint+"/vhosts", nil)
	if err != nil {
		client.logger.Error("Couldn't create request", err)
		return infos, err
	}
	req.SetBasicAuth(client.config.RMQManagement.User, client.config.RMQManagement.Password)

	res, err := client.httpClient.Do(req)
	if err != nil {
		client.logger.Error("Couldn't make request", err)
		return infos, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		client.logger.Error("Couldn't read body", err)
		return infos, err
	}

	err = yaml.Unmarshal(body, &infos)
	if err != nil {
		client.logger.Error("Couldn't deserialize response", err)
		return infos, err
	}
	return infos, nil
}

func (client *Client) GetQueues(vhost string) ([]QueueInfo, error) {
	var infos []QueueInfo

	req, err := http.NewRequest(http.MethodGet, client.config.RMQManagement.Endpoint+"/queues/"+vhost, nil)
	if err != nil {
		client.logger.Error("Couldn't create request", err)
		return infos, err
	}
	req.SetBasicAuth(client.config.RMQManagement.User, client.config.RMQManagement.Password)

	res, err := client.httpClient.Do(req)
	if err != nil {
		client.logger.Error("Couldn't make request", err)
		return infos, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		client.logger.Error("Couldn't read body", err)
		return infos, err
	}

	err = yaml.Unmarshal(body, &infos)
	if err != nil {
		client.logger.Error("Couldn't deserialize response", err)
		return infos, err
	}
	return infos, nil
}
