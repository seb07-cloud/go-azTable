//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

const (
	moduleName    = "aztables"
	moduleVersion = "v1.0.1"
)

type Enum0 string

const (
	Enum0TwoThousandNineteen0202 Enum0 = "2019-02-02"
)

// PossibleEnum0Values returns the possible values for the Enum0 const type.
func PossibleEnum0Values() []Enum0 {
	return []Enum0{
		Enum0TwoThousandNineteen0202,
	}
}

type Enum1 string

const (
	Enum1Three0 Enum1 = "3.0"
)

// PossibleEnum1Values returns the possible values for the Enum1 const type.
func PossibleEnum1Values() []Enum1 {
	return []Enum1{
		Enum1Three0,
	}
}

type Enum4 string

const (
	Enum4ACL Enum4 = "acl"
)

// PossibleEnum4Values returns the possible values for the Enum4 const type.
func PossibleEnum4Values() []Enum4 {
	return []Enum4{
		Enum4ACL,
	}
}

type Enum5 string

const (
	Enum5Service Enum5 = "service"
)

// PossibleEnum5Values returns the possible values for the Enum5 const type.
func PossibleEnum5Values() []Enum5 {
	return []Enum5{
		Enum5Service,
	}
}

type Enum6 string

const (
	Enum6Properties Enum6 = "properties"
)

// PossibleEnum6Values returns the possible values for the Enum6 const type.
func PossibleEnum6Values() []Enum6 {
	return []Enum6{
		Enum6Properties,
	}
}

type Enum7 string

const (
	Enum7Stats Enum7 = "stats"
)

// PossibleEnum7Values returns the possible values for the Enum7 const type.
func PossibleEnum7Values() []Enum7 {
	return []Enum7{
		Enum7Stats,
	}
}

// GeoReplicationStatusType - The status of the secondary location.
type GeoReplicationStatusType string

const (
	GeoReplicationStatusTypeBootstrap   GeoReplicationStatusType = "bootstrap"
	GeoReplicationStatusTypeLive        GeoReplicationStatusType = "live"
	GeoReplicationStatusTypeUnavailable GeoReplicationStatusType = "unavailable"
)

// PossibleGeoReplicationStatusTypeValues returns the possible values for the GeoReplicationStatusType const type.
func PossibleGeoReplicationStatusTypeValues() []GeoReplicationStatusType {
	return []GeoReplicationStatusType{
		GeoReplicationStatusTypeBootstrap,
		GeoReplicationStatusTypeLive,
		GeoReplicationStatusTypeUnavailable,
	}
}

type ODataMetadataFormat string

const (
	ODataMetadataFormatApplicationJSONODataFullmetadata    ODataMetadataFormat = "application/json;odata=fullmetadata"
	ODataMetadataFormatApplicationJSONODataMinimalmetadata ODataMetadataFormat = "application/json;odata=minimalmetadata"
	ODataMetadataFormatApplicationJSONODataNometadata      ODataMetadataFormat = "application/json;odata=nometadata"
)

// PossibleODataMetadataFormatValues returns the possible values for the ODataMetadataFormat const type.
func PossibleODataMetadataFormatValues() []ODataMetadataFormat {
	return []ODataMetadataFormat{
		ODataMetadataFormatApplicationJSONODataFullmetadata,
		ODataMetadataFormatApplicationJSONODataMinimalmetadata,
		ODataMetadataFormatApplicationJSONODataNometadata,
	}
}

type ResponseFormat string

const (
	ResponseFormatReturnContent   ResponseFormat = "return-content"
	ResponseFormatReturnNoContent ResponseFormat = "return-no-content"
)

// PossibleResponseFormatValues returns the possible values for the ResponseFormat const type.
func PossibleResponseFormatValues() []ResponseFormat {
	return []ResponseFormat{
		ResponseFormatReturnContent,
		ResponseFormatReturnNoContent,
	}
}