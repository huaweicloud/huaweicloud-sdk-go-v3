package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateVaultResourceResponse Response Object
type MigrateVaultResourceResponse struct {

	//
	MigratedResources *[]string `json:"migrated_resources,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o MigrateVaultResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceResponse struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceResponse", string(data)}, " ")
}
