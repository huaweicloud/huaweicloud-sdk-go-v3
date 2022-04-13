package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepoIdResponse struct {
	Error *Error `json:"error,omitempty"`
	// 响应结果

	Result *int32 `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepoIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoIdResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoIdResponse", string(data)}, " ")
}
