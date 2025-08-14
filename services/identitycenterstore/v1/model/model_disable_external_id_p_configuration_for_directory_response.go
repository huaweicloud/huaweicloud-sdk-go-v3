package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableExternalIdPConfigurationForDirectoryResponse Response Object
type DisableExternalIdPConfigurationForDirectoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableExternalIdPConfigurationForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableExternalIdPConfigurationForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"DisableExternalIdPConfigurationForDirectoryResponse", string(data)}, " ")
}
