package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthorizeHandlerRequestBody struct {

	// 数据 ID
	Ids []string `json:"ids"`

	// 0-同意 2-拒绝 3-取消
	Type int32 `json:"type"`
}

func (o AuthorizeHandlerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeHandlerRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeHandlerRequestBody", string(data)}, " ")
}
