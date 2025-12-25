package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentConfigurationResponse Response Object
type ListComponentConfigurationResponse struct {

	// 组件配置数量
	Count *int64 `json:"count,omitempty"`

	// 组件配置列表
	Records        *[]ComponentConfiguration `json:"records,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListComponentConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ListComponentConfigurationResponse", string(data)}, " ")
}
