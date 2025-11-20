package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceConfigurationsResponse Response Object
type ListResourceConfigurationsResponse struct {

	// 提权访问中的资源配置。
	ResourceConfigurations *[]ResourceConfiguration `json:"resource_configurations,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResourceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceConfigurationsResponse", string(data)}, " ")
}
