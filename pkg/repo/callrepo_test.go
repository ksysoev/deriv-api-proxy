package repo

import (
	"testing"

	"github.com/ksysoev/deriv-api-bff/pkg/core/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewCallsRepository(t *testing.T) {
	tests := []struct {
		cfg     *CallsConfig
		name    string
		wantErr bool
	}{
		{
			name: "valid config",
			cfg: &CallsConfig{
				Calls: []CallConfig{
					{
						Method: "testMethod",
						Params: validator.Config{"param1": {Type: "string"}},
						Backend: []BackendConfig{
							{
								FieldsMap:       map[string]string{"field1": "value1"},
								ResponseBody:    "responseBody1",
								RequestTemplate: "template1",
								Allow:           []string{"allow1"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid template",
			cfg: &CallsConfig{
				Calls: []CallConfig{
					{
						Method: "testMethod",
						Params: validator.Config{"param1": {Type: "value1"}},
						Backend: []BackendConfig{
							{
								FieldsMap:       map[string]string{"field1": "value1"},
								ResponseBody:    "responseBody1",
								RequestTemplate: "{{.InvalidTemplate",
								Allow:           []string{"allow1"},
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCallsRepository(tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Contains(t, got.calls, "testMethod")
			}
		})
	}
}
func TestGetCall(t *testing.T) {
	repo, err := NewCallsRepository(&CallsConfig{
		Calls: []CallConfig{
			{
				Method: "testMethod",
				Params: validator.Config{"param1": {Type: "string"}},
				Backend: []BackendConfig{
					{
						FieldsMap:       map[string]string{"field1": "value1"},
						ResponseBody:    "responseBody1",
						RequestTemplate: "template1",
						Allow:           []string{"allow1"},
					},
				},
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, repo)

	tests := []struct {
		name   string
		method string
		found  bool
	}{
		{
			name:   "existing method",
			method: "testMethod",
			found:  true,
		},
		{
			name:   "non-existing method",
			method: "nonExistingMethod",
			found:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callConfig := repo.GetCall(tt.method)

			if tt.found {
				assert.NotNil(t, callConfig)
			} else {
				assert.Nil(t, callConfig)
			}
		})
	}
}
func TestSortBackends(t *testing.T) {
	tests := []struct {
		name    string
		input   []BackendConfig
		want    []BackendConfig
		wantErr bool
	}{
		{
			name: "no dependencies",
			input: []BackendConfig{
				{ResponseBody: "response1"},
				{ResponseBody: "response2"},
			},
			want: []BackendConfig{
				{ResponseBody: "response1"},
				{ResponseBody: "response2"},
			},
			wantErr: false,
		},
		{
			name: "simple dependency",
			input: []BackendConfig{
				{ResponseBody: "response1", DependsOn: []string{"response2"}},
				{ResponseBody: "response2"},
			},
			want: []BackendConfig{
				{ResponseBody: "response2"},
				{ResponseBody: "response1", DependsOn: []string{"response2"}},
			},
			wantErr: false,
		},
		{
			name: "circular dependency",
			input: []BackendConfig{
				{ResponseBody: "response1", DependsOn: []string{"response2"}},
				{ResponseBody: "response2", DependsOn: []string{"response1"}},
			},
			want:    nil,
			wantErr: true,
		},

		//TODO: Fix the test case
		{
			name: "complex dependency",
			input: []BackendConfig{
				{ResponseBody: "response1", DependsOn: []string{"response3"}},
				{ResponseBody: "response2", DependsOn: []string{"response1"}},
				{ResponseBody: "response3"},
			},
			want: []BackendConfig{
				{ResponseBody: "response3"},
				{ResponseBody: "response1", DependsOn: []string{"response3"}},
				{ResponseBody: "response2", DependsOn: []string{"response1"}},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortBackends(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
