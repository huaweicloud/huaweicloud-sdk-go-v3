package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentConfigurationsResponse Response Object
type ListComponentConfigurationsResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“ComponentConfiguration”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// 组件配置列表。
	Items          *[]Configuration `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListComponentConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentConfigurationsResponse", string(data)}, " ")
}
