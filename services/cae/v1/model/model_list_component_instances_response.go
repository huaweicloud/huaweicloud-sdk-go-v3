package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentInstancesResponse Response Object
type ListComponentInstancesResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“ComponentInstance”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
