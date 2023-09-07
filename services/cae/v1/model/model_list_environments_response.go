package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentsResponse Response Object
type ListEnvironmentsResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *EnvironmentKindObj `json:"kind,omitempty"`

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
