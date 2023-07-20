package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsResponse Response Object
type ListConfigurationsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 用户可创建的自定义参数模板最大数量。
	Quota *int32 `json:"quota,omitempty"`

	Configurations *[]ListConfigurationsResult `json:"configurations,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}
