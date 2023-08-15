package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunModifyPictureResponse Response Object
type RunModifyPictureResponse struct {

	// 调用成功时表示调用结果。  调用失败时无此字段。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunModifyPictureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModifyPictureResponse struct{}"
	}

	return strings.Join([]string{"RunModifyPictureResponse", string(data)}, " ")
}
