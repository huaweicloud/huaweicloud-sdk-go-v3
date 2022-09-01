package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepoIdResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 响应结果
	Result *int32 `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepoIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoIdResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoIdResponse", string(data)}, " ")
}
