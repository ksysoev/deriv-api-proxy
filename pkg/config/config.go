package config

import (
	"reflect"

	"github.com/ksysoev/deriv-api-bff/pkg/api"
	"github.com/ksysoev/deriv-api-bff/pkg/core/validator"
	"github.com/ksysoev/deriv-api-bff/pkg/prov/deriv"
)

type Config struct {
	Server api.Config   `mapstructure:"server"`
	Deriv  deriv.Config `mapstructure:"deriv"`
	API    CallsConfig  `mapstructure:"api"`
	Etcd   EtcdConfig   `mapstructure:"etcd"`
}

type CallsConfig struct {
	Calls []CallConfig `mapstructure:"calls"`
}

type EtcdConfig struct {
	Servers            []string `mapstructure:"servers"`
	DialTimeoutSeconds int      `mapstructure:"dialTimeoutSeconds"`
}

type CallConfig struct {
	Method  string           `mapstructure:"method"`
	Params  validator.Config `mapstructure:"params"`
	Backend []*BackendConfig `mapstructure:"backend"`
}

type BackendConfig struct {
	Name            string            `mapstructure:"name"`
	FieldsMap       map[string]string `mapstructure:"fields_map"`
	ResponseBody    string            `mapstructure:"response_body"`
	RequestTemplate map[string]any    `mapstructure:"request_template"`
	Method          string            `mapstructure:"method"`
	URLTemplate     string            `mapstructure:"url_template"`
	DependsOn       []string          `mapstructure:"depends_on"`
	Allow           []string          `mapstructure:"allow"`
}

// TODO: add godoc
func Compare(_old, _new interface{}, path string) []string {
	var diffs []string

	oldMeta := reflect.ValueOf(_old)
	newMeta := reflect.ValueOf(_new)

	for i := 0; i < oldMeta.NumField(); i++ {
		oldField := oldMeta.Field(i)
		newField := newMeta.Field(i)
		fieldName := oldMeta.Type().Field(i).Name

		currentPath := path + "." + fieldName

		if oldField.Kind() == reflect.Struct {
			nestedDiffs := Compare(oldField.Interface(), newField.Interface(), currentPath)
			diffs = append(diffs, nestedDiffs...)
		} else if !reflect.DeepEqual(oldField.Interface(), newField.Interface()) {
			diffs = append(diffs, currentPath)
		}
	}

	return diffs
}
