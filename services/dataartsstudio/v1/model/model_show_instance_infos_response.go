package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceInfosResponse Response Object
type ShowInstanceInfosResponse struct {

	// 响应码
	StatusCode *int32 `json:"status_code,omitempty"`

	// 响应体
	Body *string `json:"body,omitempty"`

	// 响应头，结构为Map<String,String>
	HeaderMap      *interface{} `json:"header_map,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowInstanceInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceInfosResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceInfosResponse", string(data)}, " ")
}
