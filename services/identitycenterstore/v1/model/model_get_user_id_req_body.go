package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetUserIdReqBody 获取用户id的请求体
type GetUserIdReqBody struct {
	AlternateIdentifier *AlternateIdentifierDto `json:"alternate_identifier"`
}

func (o GetUserIdReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUserIdReqBody struct{}"
	}

	return strings.Join([]string{"GetUserIdReqBody", string(data)}, " ")
}
