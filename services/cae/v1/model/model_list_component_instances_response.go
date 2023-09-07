package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentInstancesResponse Response Object
type ListComponentInstancesResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentConfigurationKindObj `json:"kind,omitempty"`

	// 组件实例列表。
	Items          *[]Instance `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListComponentInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListComponentInstancesResponse", string(data)}, " ")
}
