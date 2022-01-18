package options

import (
	"encoding/json"
	"reflect"
	"testing"
)

const (
	postgresqlOptions = `{
		"dbname": "aa",
		"host": "1.1.1.1",
		"port": 5432
	}`
	compactPostgresqlOptions = `{"dbname": "aa","host": "1.1.1.1","port": 5432}`
	queryOptions             = `{
		"parameters": []
	}`
	compactQueryOptions = `{"parameters": []}`
	invalidQueryOptions = `{
		"parameters": [],
	}`
)

func TestMapFromString(t *testing.T) {
	var dataSourceOptions map[string]interface{}
	_ = json.Unmarshal([]byte(compactPostgresqlOptions), &dataSourceOptions)
	var queryOptionsMap map[string]interface{}
	_ = json.Unmarshal([]byte(compactQueryOptions), &queryOptionsMap)
	type args struct {
		optionsJSON string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "valid data source json",
			args:    args{optionsJSON: postgresqlOptions},
			want:    dataSourceOptions,
			wantErr: false,
		},
		{
			name:    "valid query json",
			args:    args{optionsJSON: queryOptions},
			want:    queryOptionsMap,
			wantErr: false,
		},
		{
			name:    "empty json",
			args:    args{optionsJSON: "{}"},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name:    "invalid json",
			args:    args{optionsJSON: invalidQueryOptions},
			want:    map[string]interface{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapFromString(tt.args.optionsJSON)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
