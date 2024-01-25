package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttp2RpcsResponse Response Object
type ShowHttp2RpcsResponse struct {

	// http2Rpc 总数。
	Total *int32 `json:"total,omitempty"`

	// Http2Rpc 详细信息。
	Data           *[]Http2Rpc `json:"data,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowHttp2RpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttp2RpcsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttp2RpcsResponse", string(data)}, " ")
}
