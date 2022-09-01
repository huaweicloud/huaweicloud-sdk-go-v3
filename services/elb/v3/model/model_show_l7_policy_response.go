package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowL7PolicyResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	L7policy       *L7Policy `json:"l7policy,omitempty" xml:"l7policy"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowL7PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowL7PolicyResponse", string(data)}, " ")
}
