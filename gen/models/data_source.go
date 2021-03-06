// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataSource data source
//
// swagger:model dataSource
type DataSource struct {

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// options
	// Required: true
	Options interface{} `json:"options"`

	// syntax
	// Min Length: 1
	Syntax string `json:"syntax,omitempty"`

	// type
	// Required: true
	// Enum: [aws_es athena azure_kusto bigquery_gce Cassandra scylla clickhouse cloudwatch_insights corporate_memory databricks db2 dgraph dynamodb_sql google_analytics google_spreadsheets hive hive_http impala influxdb memsql mssql mssql_odbc rds_mysql oracle pg redshift redshift_iam cockroach phoenix presto qubole rockset insecure_script sparql_endpoint sqlite treasuredata trino yandex_metrika yandex_appmetrika]
	Type *string `json:"type"`
}

// Validate validates this data source
func (m *DataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyntax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateOptions(formats strfmt.Registry) error {

	if m.Options == nil {
		return errors.Required("options", "body", nil)
	}

	return nil
}

func (m *DataSource) validateSyntax(formats strfmt.Registry) error {
	if swag.IsZero(m.Syntax) { // not required
		return nil
	}

	if err := validate.MinLength("syntax", "body", m.Syntax, 1); err != nil {
		return err
	}

	return nil
}

var dataSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aws_es","athena","azure_kusto","bigquery_gce","Cassandra","scylla","clickhouse","cloudwatch_insights","corporate_memory","databricks","db2","dgraph","dynamodb_sql","google_analytics","google_spreadsheets","hive","hive_http","impala","influxdb","memsql","mssql","mssql_odbc","rds_mysql","oracle","pg","redshift","redshift_iam","cockroach","phoenix","presto","qubole","rockset","insecure_script","sparql_endpoint","sqlite","treasuredata","trino","yandex_metrika","yandex_appmetrika"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataSourceTypeTypePropEnum = append(dataSourceTypeTypePropEnum, v)
	}
}

const (

	// DataSourceTypeAwsEs captures enum value "aws_es"
	DataSourceTypeAwsEs string = "aws_es"

	// DataSourceTypeAthena captures enum value "athena"
	DataSourceTypeAthena string = "athena"

	// DataSourceTypeAzureKusto captures enum value "azure_kusto"
	DataSourceTypeAzureKusto string = "azure_kusto"

	// DataSourceTypeBigqueryGce captures enum value "bigquery_gce"
	DataSourceTypeBigqueryGce string = "bigquery_gce"

	// DataSourceTypeCassandra captures enum value "Cassandra"
	DataSourceTypeCassandra string = "Cassandra"

	// DataSourceTypeScylla captures enum value "scylla"
	DataSourceTypeScylla string = "scylla"

	// DataSourceTypeClickhouse captures enum value "clickhouse"
	DataSourceTypeClickhouse string = "clickhouse"

	// DataSourceTypeCloudwatchInsights captures enum value "cloudwatch_insights"
	DataSourceTypeCloudwatchInsights string = "cloudwatch_insights"

	// DataSourceTypeCorporateMemory captures enum value "corporate_memory"
	DataSourceTypeCorporateMemory string = "corporate_memory"

	// DataSourceTypeDatabricks captures enum value "databricks"
	DataSourceTypeDatabricks string = "databricks"

	// DataSourceTypeDb2 captures enum value "db2"
	DataSourceTypeDb2 string = "db2"

	// DataSourceTypeDgraph captures enum value "dgraph"
	DataSourceTypeDgraph string = "dgraph"

	// DataSourceTypeDynamodbSQL captures enum value "dynamodb_sql"
	DataSourceTypeDynamodbSQL string = "dynamodb_sql"

	// DataSourceTypeGoogleAnalytics captures enum value "google_analytics"
	DataSourceTypeGoogleAnalytics string = "google_analytics"

	// DataSourceTypeGoogleSpreadsheets captures enum value "google_spreadsheets"
	DataSourceTypeGoogleSpreadsheets string = "google_spreadsheets"

	// DataSourceTypeHive captures enum value "hive"
	DataSourceTypeHive string = "hive"

	// DataSourceTypeHiveHTTP captures enum value "hive_http"
	DataSourceTypeHiveHTTP string = "hive_http"

	// DataSourceTypeImpala captures enum value "impala"
	DataSourceTypeImpala string = "impala"

	// DataSourceTypeInfluxdb captures enum value "influxdb"
	DataSourceTypeInfluxdb string = "influxdb"

	// DataSourceTypeMemsql captures enum value "memsql"
	DataSourceTypeMemsql string = "memsql"

	// DataSourceTypeMssql captures enum value "mssql"
	DataSourceTypeMssql string = "mssql"

	// DataSourceTypeMssqlOdbc captures enum value "mssql_odbc"
	DataSourceTypeMssqlOdbc string = "mssql_odbc"

	// DataSourceTypeRdsMysql captures enum value "rds_mysql"
	DataSourceTypeRdsMysql string = "rds_mysql"

	// DataSourceTypeOracle captures enum value "oracle"
	DataSourceTypeOracle string = "oracle"

	// DataSourceTypePg captures enum value "pg"
	DataSourceTypePg string = "pg"

	// DataSourceTypeRedshift captures enum value "redshift"
	DataSourceTypeRedshift string = "redshift"

	// DataSourceTypeRedshiftIam captures enum value "redshift_iam"
	DataSourceTypeRedshiftIam string = "redshift_iam"

	// DataSourceTypeCockroach captures enum value "cockroach"
	DataSourceTypeCockroach string = "cockroach"

	// DataSourceTypePhoenix captures enum value "phoenix"
	DataSourceTypePhoenix string = "phoenix"

	// DataSourceTypePresto captures enum value "presto"
	DataSourceTypePresto string = "presto"

	// DataSourceTypeQubole captures enum value "qubole"
	DataSourceTypeQubole string = "qubole"

	// DataSourceTypeRockset captures enum value "rockset"
	DataSourceTypeRockset string = "rockset"

	// DataSourceTypeInsecureScript captures enum value "insecure_script"
	DataSourceTypeInsecureScript string = "insecure_script"

	// DataSourceTypeSparqlEndpoint captures enum value "sparql_endpoint"
	DataSourceTypeSparqlEndpoint string = "sparql_endpoint"

	// DataSourceTypeSqlite captures enum value "sqlite"
	DataSourceTypeSqlite string = "sqlite"

	// DataSourceTypeTreasuredata captures enum value "treasuredata"
	DataSourceTypeTreasuredata string = "treasuredata"

	// DataSourceTypeTrino captures enum value "trino"
	DataSourceTypeTrino string = "trino"

	// DataSourceTypeYandexMetrika captures enum value "yandex_metrika"
	DataSourceTypeYandexMetrika string = "yandex_metrika"

	// DataSourceTypeYandexAppmetrika captures enum value "yandex_appmetrika"
	DataSourceTypeYandexAppmetrika string = "yandex_appmetrika"
)

// prop value enum
func (m *DataSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataSource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this data source based on the context it is used
func (m *DataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSource) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSource) UnmarshalBinary(b []byte) error {
	var res DataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
