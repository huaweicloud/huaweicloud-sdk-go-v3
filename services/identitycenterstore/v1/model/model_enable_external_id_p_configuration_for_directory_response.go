package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableExternalIdPConfigurationForDirectoryResponse Response Object
type EnableExternalIdPConfigurationForDirectoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableExternalIdPConfigurationForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableExternalIdPConfigurationForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"EnableExternalIdPConfigurationForDirectoryResponse", string(data)}, " ")
}
