package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DiscoverEventSchemaFromDataResponse struct {

	// 自动发现后的模型
	Definition *string `json:"definition,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DiscoverEventSchemaFromDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoverEventSchemaFromDataResponse struct{}"
	}

	return strings.Join([]string{"DiscoverEventSchemaFromDataResponse", string(data)}, " ")
}
