package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SamlConfigurationPropertyItemsString represents the SamlConfigurationPropertyItemsString schema from the OpenAPI specification
type SamlConfigurationPropertyItemsString struct {
	Optional bool `json:"optional,omitempty"` // True if optional
	TypeField int `json:"type,omitempty"` // Property type, 1=String, 3=long, 11=boolean, 12=Password
	Value string `json:"value,omitempty"` // Property value
	Description string `json:"description,omitempty"` // Property description
	Is_set bool `json:"is_set,omitempty"` // True if property is set
	Name string `json:"name,omitempty"` // property name
}

// TruststoreItems represents the TruststoreItems schema from the OpenAPI specification
type TruststoreItems struct {
	Alias string `json:"alias,omitempty"` // Truststore alias name
	Entrytype string `json:"entryType,omitempty"`
	Issuer string `json:"issuer,omitempty"` // e.g. "CN=Admin"
	Notafter string `json:"notAfter,omitempty"` // e.g. "Sun Jun 30 23:59:50 AEST 2019"
	Notbefore string `json:"notBefore,omitempty"` // e.g. "Sun Jul 01 12:00:00 AEST 2018"
	Serialnumber int `json:"serialNumber,omitempty"` // 18165099476682912368
	Subject string `json:"subject,omitempty"` // e.g. "CN=localhost"
}

// BundleDataProp represents the BundleDataProp schema from the OpenAPI specification
type BundleDataProp struct {
	Key string `json:"key,omitempty"` // Bundle data key
	Value string `json:"value,omitempty"` // Bundle data value
}

// BundleData represents the BundleData schema from the OpenAPI specification
type BundleData struct {
	Stateraw int `json:"stateRaw,omitempty"` // Numeric raw bundle state value
	Version string `json:"version,omitempty"` // Bundle version
	Props []BundleDataProp `json:"props,omitempty"`
	Id int `json:"id,omitempty"` // Bundle ID
	State string `json:"state,omitempty"` // Bundle state value
	Name string `json:"name,omitempty"` // Bundle name
	Symbolicname string `json:"symbolicName,omitempty"` // Bundle symbolic name
	Category string `json:"category,omitempty"` // Bundle category
	Fragment bool `json:"fragment,omitempty"` // Is bundle a fragment
}

// KeystoreChainItems represents the KeystoreChainItems schema from the OpenAPI specification
type KeystoreChainItems struct {
	Serialnumber int `json:"serialNumber,omitempty"` // 18165099476682912368
	Subject string `json:"subject,omitempty"` // e.g. "CN=localhost"
	Issuer string `json:"issuer,omitempty"` // e.g. "CN=Admin"
	Notafter string `json:"notAfter,omitempty"` // e.g. "Sun Jun 30 23:59:50 AEST 2019"
	Notbefore string `json:"notBefore,omitempty"` // e.g. "Sun Jul 01 12:00:00 AEST 2018"
}

// SamlConfigurationInfo represents the SamlConfigurationInfo schema from the OpenAPI specification
type SamlConfigurationInfo struct {
	Service_location string `json:"service_location,omitempty"` // needed for configuraiton binding
	Title string `json:"title,omitempty"` // Title
	Bundle_location string `json:"bundle_location,omitempty"` // needed for configuration binding
	Description string `json:"description,omitempty"` // Title
	Pid string `json:"pid,omitempty"` // Persistent Identity (PID)
	Properties SamlConfigurationProperties `json:"properties,omitempty"`
}

// SamlConfigurationPropertyItemsArray represents the SamlConfigurationPropertyItemsArray schema from the OpenAPI specification
type SamlConfigurationPropertyItemsArray struct {
	Name string `json:"name,omitempty"` // property name
	Optional bool `json:"optional,omitempty"` // True if optional
	TypeField int `json:"type,omitempty"` // Property type, 1=String, 3=long, 11=boolean, 12=Password
	Values []string `json:"values,omitempty"` // Property value
	Description string `json:"description,omitempty"` // Property description
	Is_set bool `json:"is_set,omitempty"` // True if property is set
}

// SamlConfigurationPropertyItemsBoolean represents the SamlConfigurationPropertyItemsBoolean schema from the OpenAPI specification
type SamlConfigurationPropertyItemsBoolean struct {
	Is_set bool `json:"is_set,omitempty"` // True if property is set
	Name string `json:"name,omitempty"` // property name
	Optional bool `json:"optional,omitempty"` // True if optional
	TypeField int `json:"type,omitempty"` // Property type, 1=String, 3=long, 11=boolean, 12=Password
	Value bool `json:"value,omitempty"` // Property value
	Description string `json:"description,omitempty"` // Property description
}

// KeystoreInfo represents the KeystoreInfo schema from the OpenAPI specification
type KeystoreInfo struct {
	Aliases []KeystoreItems `json:"aliases,omitempty"`
	Exists bool `json:"exists,omitempty"` // False if truststore don't exist
}

// SamlConfigurationProperties represents the SamlConfigurationProperties schema from the OpenAPI specification
type SamlConfigurationProperties struct {
	Idphttpredirect SamlConfigurationPropertyItemsBoolean `json:"idpHttpRedirect,omitempty"`
	Path SamlConfigurationPropertyItemsArray `json:"path,omitempty"`
	Signaturemethod SamlConfigurationPropertyItemsString `json:"signatureMethod,omitempty"`
	Serviceproviderentityid SamlConfigurationPropertyItemsString `json:"serviceProviderEntityId,omitempty"`
	Defaultredirecturl SamlConfigurationPropertyItemsString `json:"defaultRedirectUrl,omitempty"`
	Keystorepassword SamlConfigurationPropertyItemsString `json:"keyStorePassword,omitempty"`
	Handlelogout SamlConfigurationPropertyItemsBoolean `json:"handleLogout,omitempty"`
	Idpcertalias SamlConfigurationPropertyItemsString `json:"idpCertAlias,omitempty"`
	Assertionconsumerserviceurl SamlConfigurationPropertyItemsString `json:"assertionConsumerServiceURL,omitempty"`
	Useridattribute SamlConfigurationPropertyItemsString `json:"userIDAttribute,omitempty"`
	Nameidformat SamlConfigurationPropertyItemsString `json:"nameIdFormat,omitempty"`
	Useencryption SamlConfigurationPropertyItemsBoolean `json:"useEncryption,omitempty"`
	Logouturl SamlConfigurationPropertyItemsString `json:"logoutUrl,omitempty"`
	Clocktolerance SamlConfigurationPropertyItemsLong `json:"clockTolerance,omitempty"`
	Groupmembershipattribute SamlConfigurationPropertyItemsString `json:"groupMembershipAttribute,omitempty"`
	Service_ranking SamlConfigurationPropertyItemsLong `json:"service.ranking,omitempty"`
	Idpurl SamlConfigurationPropertyItemsString `json:"idpUrl,omitempty"`
	Defaultgroups SamlConfigurationPropertyItemsArray `json:"defaultGroups,omitempty"`
	Createuser SamlConfigurationPropertyItemsBoolean `json:"createUser,omitempty"`
	Spprivatekeyalias SamlConfigurationPropertyItemsString `json:"spPrivateKeyAlias,omitempty"`
	Synchronizeattributes SamlConfigurationPropertyItemsArray `json:"synchronizeAttributes,omitempty"`
	Addgroupmemberships SamlConfigurationPropertyItemsBoolean `json:"addGroupMemberships,omitempty"`
	Digestmethod SamlConfigurationPropertyItemsString `json:"digestMethod,omitempty"`
	Userintermediatepath SamlConfigurationPropertyItemsString `json:"userIntermediatePath,omitempty"`
}

// BundleInfo represents the BundleInfo schema from the OpenAPI specification
type BundleInfo struct {
	S []int `json:"s,omitempty"`
	Status string `json:"status,omitempty"` // Status description of all bundles
	Data []BundleData `json:"data,omitempty"`
}

// InstallStatus represents the InstallStatus schema from the OpenAPI specification
type InstallStatus struct {
	Status map[string]interface{} `json:"status,omitempty"`
}

// KeystoreItems represents the KeystoreItems schema from the OpenAPI specification
type KeystoreItems struct {
	Format string `json:"format,omitempty"` // e.g. "PKCS#8"
	Algorithm string `json:"algorithm,omitempty"` // e.g. "RSA"
	Alias string `json:"alias,omitempty"` // Keystore alias name
	Chain []KeystoreChainItems `json:"chain,omitempty"`
	Entrytype string `json:"entryType,omitempty"` // e.g. "privateKey"
}

// TruststoreInfo represents the TruststoreInfo schema from the OpenAPI specification
type TruststoreInfo struct {
	Aliases []TruststoreItems `json:"aliases,omitempty"`
	Exists bool `json:"exists,omitempty"` // False if truststore don't exist
}

// SamlConfigurationPropertyItemsLong represents the SamlConfigurationPropertyItemsLong schema from the OpenAPI specification
type SamlConfigurationPropertyItemsLong struct {
	Description string `json:"description,omitempty"` // Property description
	Is_set bool `json:"is_set,omitempty"` // True if property is set
	Name string `json:"name,omitempty"` // property name
	Optional bool `json:"optional,omitempty"` // True if optional
	TypeField int `json:"type,omitempty"` // Property type, 1=String, 3=long, 11=boolean, 12=Password
	Value int `json:"value,omitempty"` // Property value
}
