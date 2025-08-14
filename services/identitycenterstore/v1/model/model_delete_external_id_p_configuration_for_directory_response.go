package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExternalIdPConfigurationForDirectoryResponse Response Object
type DeleteExternalIdPConfigurationForDirectoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExternalIdPConfigurationForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExternalIdPConfigurationForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteExternalIdPConfigurationForDirectoryResponse", string(data)}, " ")
}
