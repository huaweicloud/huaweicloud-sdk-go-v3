package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsResponse Response Object
type ListConfigurationsResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 参数模板列表。
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
