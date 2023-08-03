module github.com/starkandwayne/rabbitmq-metrics-emitter

go 1.14

replace github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200531234253-77e06fda0c94+incompatible

replace github.com/SermoDigital/jose => github.com/SermoDigital/jose v0.9.2-0.20161205224733-f6df55f235c2

replace github.com/mailru/easyjson => github.com/mailru/easyjson v0.0.0-20180323154445-8b799c424f57

replace github.com/cloudfoundry/sonde-go => github.com/cloudfoundry/sonde-go v0.0.0-20171206171820-b33733203bb4

replace code.cloudfoundry.org/go-log-cache => code.cloudfoundry.org/go-log-cache v1.0.1-0.20200316170138-f466e0302c34

require (
	code.cloudfoundry.org/cli v7.1.0+incompatible
	code.cloudfoundry.org/diego-ssh v0.0.0-20230612151408-7461829a983b // indirect
	code.cloudfoundry.org/go-diodes v0.0.0-20190809170250-f77fb823c7ee // indirect
	code.cloudfoundry.org/go-log-cache v0.0.0-00010101000000-000000000000 // indirect
	code.cloudfoundry.org/go-loggregator v7.4.0+incompatible
	code.cloudfoundry.org/inigo v0.0.0-20230612153013-b300679e6ed6 // indirect
	code.cloudfoundry.org/jsonry v1.1.0 // indirect
	code.cloudfoundry.org/lager v2.0.0+incompatible
	code.cloudfoundry.org/lager/v3 v3.0.2 // indirect
	github.com/cloudfoundry-community/go-cf-clients-helper v1.0.1
	github.com/moby/moby v20.10.25+incompatible // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/onsi/ginkgo/v2 v2.11.0 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/tedsuo/ifrit v0.0.0-20230516164442-7862c310ad26 // indirect
	gopkg.in/yaml.v3 v3.0.1
)
