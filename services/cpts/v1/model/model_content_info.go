package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContentInfo struct {

	// body_type
	BodyType *int32 `json:"body_type,omitempty" xml:"body_type"`

	// bodys
	Bodys *[]interface{} `json:"bodys,omitempty" xml:"bodys"`

	// check_end_length
	CheckEndLength *string `json:"check_end_length,omitempty" xml:"check_end_length"`

	// check_end_str
	CheckEndStr *string `json:"check_end_str,omitempty" xml:"check_end_str"`

	// check_end_type
	CheckEndType *string `json:"check_end_type,omitempty" xml:"check_end_type"`

	// connect_timeout
	ConnectTimeout *int32 `json:"connect_timeout,omitempty" xml:"connect_timeout"`

	// connect_type
	ConnectType *int32 `json:"connect_type,omitempty" xml:"connect_type"`

	// headers
	Headers *[]ContentHeader `json:"headers,omitempty" xml:"headers"`

	// http_version
	HttpVersion *string `json:"http_version,omitempty" xml:"http_version"`

	// method
	Method *string `json:"method,omitempty" xml:"method"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// protocol_type
	ProtocolType *int32 `json:"protocol_type,omitempty" xml:"protocol_type"`

	// return_timeout
	ReturnTimeout *int32 `json:"return_timeout,omitempty" xml:"return_timeout"`

	// return_timeout_param
	ReturnTimeoutParam *string `json:"return_timeout_param,omitempty" xml:"return_timeout_param"`

	// url
	Url *string `json:"url,omitempty" xml:"url"`
}

func (o ContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentInfo struct{}"
	}

	return strings.Join([]string{"ContentInfo", string(data)}, " ")
}
