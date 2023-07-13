package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentsResponse Response Object
type ListEnvironmentsResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Environment”，该值不可修改。
	Kind *string `json:"Kind,omitempty"`

	// 环境列表。
	Items          *[]EnvironmentItem `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListEnvironmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsResponse", string(data)}, " ")
}
