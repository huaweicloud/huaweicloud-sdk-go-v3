package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunAddPictureResponse struct {
	// 调用成功时表示调用结果。  调用失败时无此字段。

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunAddPictureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAddPictureResponse struct{}"
	}

	return strings.Join([]string{"RunAddPictureResponse", string(data)}, " ")
}
