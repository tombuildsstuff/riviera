package storage

import "github.com/jen20/riviera/azure"

type CreateStorageAccountResponse struct {
	Location    			*string `mapstructure:"location"`
	AccountType 			*string `mapstructure:"accountType"`
	EncryptionEnabled *string `mapstructure:"encryptionEnabled"`
	AccessTier 				bool `mapstructure:"accessTier"`
}

type CreateStorageAccount struct {
	Name              string             `json:"-"`
	ResourceGroupName string             `json:"-"`
	AccountType       *string            `json:"accountType,omitempty"`
	Location          string             `json:"-" riviera:"location"`
	Tags              map[string]*string `json:"-" riviera:"tags"`
	EncryptionEnabled bool							 `json:"-"`
	AccessTier				string						 `json:"-"`
}

func (s CreateStorageAccount) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "PUT",
		URLPathFunc: storageDefaultURLPathFunc(s.ResourceGroupName, s.Name),
		ResponseTypeFunc: func() interface{} {
			return &CreateStorageAccountResponse{}
		},
	}
}
