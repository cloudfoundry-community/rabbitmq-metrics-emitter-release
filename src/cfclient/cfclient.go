package cfclient

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	"github.com/cloudfoundry-community/go-cf-clients-helper"
	"github.com/starkandwayne/rabbit-mq-metrics-emitter/config"
)

type CfClient struct {
	inner *ccv2.Client
}

func NewCfClient(config *config.Config) (*CfClient, error) {

	clientConfig := clients.Config{
		Endpoint:          config.CfConfig.ApiUrl,
		SkipSslValidation: config.CfConfig.SkipSslValidation,
		User:              config.CfConfig.Username,
		Password:          config.CfConfig.Password,
	}

	session, err := clients.NewSession(clientConfig)
	if err != nil {
		return nil, err
	}
	return &CfClient{
		inner: session.V2()}, err
}

func (client *CfClient) AllBindings(instanceIds []string) ([]ccv2.ServiceBinding, error) {
	bindings, _, err := client.inner.GetServiceBindings(ccv2.Filter{
		Type:     constant.ServiceInstanceGUIDFilter,
		Operator: constant.InOperator,
		Values:   instanceIds,
	})
	return bindings, err
}
