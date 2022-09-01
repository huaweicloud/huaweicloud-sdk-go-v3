package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartInstanceResponse struct {

	// 返回值
	Result *string `json:"result,omitempty" xml:"result"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceResponse", string(data)}, " ")
}
