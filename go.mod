module get.porter.sh/mixin/docker-compose

go 1.13

require (
	get.porter.sh/porter v0.38.1
	github.com/ghodss/yaml v1.0.0
	github.com/gobuffalo/packr/v2 v2.8.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/xeipuuv/gojsonschema v1.2.0
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
