//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import "encoding/json"

func unmarshalBackupPolicyClassification(rawMsg json.RawMessage) (BackupPolicyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupPolicyClassification
	switch m["type"] {
	case string(BackupPolicyTypeContinuous):
		b = &ContinuousModeBackupPolicy{}
	case string(BackupPolicyTypePeriodic):
		b = &PeriodicModeBackupPolicy{}
	default:
		b = &BackupPolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataTransferDataSourceSinkClassification(rawMsg json.RawMessage) (DataTransferDataSourceSinkClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataTransferDataSourceSinkClassification
	switch m["component"] {
	case "BaseCosmosDataTransferDataSourceSink":
		b = &BaseCosmosDataTransferDataSourceSink{}
	case string(DataTransferComponentAzureBlobStorage):
		b = &AzureBlobDataTransferDataSourceSink{}
	case string(DataTransferComponentCosmosDBCassandra):
		b = &CassandraDataTransferDataSourceSink{}
	case string(DataTransferComponentCosmosDBMongo):
		b = &MongoDataTransferDataSourceSink{}
	case string(DataTransferComponentCosmosDBSQL):
		b = &SQLDataTransferDataSourceSink{}
	default:
		b = &DataTransferDataSourceSink{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalServiceResourcePropertiesClassification(rawMsg json.RawMessage) (ServiceResourcePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServiceResourcePropertiesClassification
	switch m["serviceType"] {
	case string(ServiceTypeDataTransfer):
		b = &DataTransferServiceResourceProperties{}
	case string(ServiceTypeGraphAPICompute):
		b = &GraphAPIComputeServiceResourceProperties{}
	case string(ServiceTypeMaterializedViewsBuilder):
		b = &MaterializedViewsBuilderServiceResourceProperties{}
	case string(ServiceTypeSQLDedicatedGateway):
		b = &SQLDedicatedGatewayServiceResourceProperties{}
	default:
		b = &ServiceResourceProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
