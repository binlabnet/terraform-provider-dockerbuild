module gitlab.com/noname-ltd/terraform-provider-dockerbuild

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Sirupsen/logrus v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.2-0.20170601211448-f5ec1e2936dc
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/pkg/errors v0.8.1
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191011234655-491137f69257 // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
)

// replace github.com/sirupsen/logrus => github.com/Sirupsen/logrus v1.4.2

replace github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2

// replace github.com/docker/docker => github.com/docker/docker v1.13.2-0.20170601211448-f5ec1e2936dc
