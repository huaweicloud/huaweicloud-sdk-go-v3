package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetGroupIdReqBody 获取用户组id请求体
type GetGroupIdReqBody struct {
	AlternateIdentifier *AlternateIdentifierDto `json:"alternate_identifier"`
}

func (o GetGroupIdReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetGroupIdReqBody struct{}"
	}

	return strings.Join([]string{"GetGroupIdReqBody", string(data)}, " ")
}
