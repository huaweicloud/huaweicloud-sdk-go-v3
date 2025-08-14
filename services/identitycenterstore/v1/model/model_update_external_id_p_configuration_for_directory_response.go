package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExternalIdPConfigurationForDirectoryResponse Response Object
type UpdateExternalIdPConfigurationForDirectoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateExternalIdPConfigurationForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExternalIdPConfigurationForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"UpdateExternalIdPConfigurationForDirectoryResponse", string(data)}, " ")
}
