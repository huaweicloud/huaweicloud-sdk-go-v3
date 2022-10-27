package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnvironmentsResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *string `json:"Kind,omitempty"`

	// 环境信息列表。
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
