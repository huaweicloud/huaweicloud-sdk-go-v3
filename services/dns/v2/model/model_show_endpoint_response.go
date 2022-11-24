package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEndpointResponse struct {

	// 查询endpoint响应。
	Endpoint       *interface{} `json:"endpoint,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointResponse struct{}"
	}

	return strings.Join([]string{"ShowEndpointResponse", string(data)}, " ")
}
