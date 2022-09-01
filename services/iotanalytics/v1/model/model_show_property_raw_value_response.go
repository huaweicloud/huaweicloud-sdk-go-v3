package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPropertyRawValueResponse struct {

	// 时间序列,使用UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z',示例：2021-02-01T00:00:00.123Z
	Timestamps *[]string `json:"timestamps,omitempty" xml:"timestamps"`

	// 响应属性列表
	Properties     *[]RawValue `json:"properties,omitempty" xml:"properties"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPropertyRawValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPropertyRawValueResponse struct{}"
	}

	return strings.Join([]string{"ShowPropertyRawValueResponse", string(data)}, " ")
}
