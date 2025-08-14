package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalIdPConfigurationsForDirectoryResponse Response Object
type ListExternalIdPConfigurationsForDirectoryResponse struct {

	// 外部身份提供商配置信息
	Associations   *[]ExternalIdpConfigurationDto `json:"associations,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListExternalIdPConfigurationsForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalIdPConfigurationsForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"ListExternalIdPConfigurationsForDirectoryResponse", string(data)}, " ")
}
