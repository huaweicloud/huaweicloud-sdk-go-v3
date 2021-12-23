package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunDeleteInstanceResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDeleteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"RunDeleteInstanceResponse", string(data)}, " ")
}
