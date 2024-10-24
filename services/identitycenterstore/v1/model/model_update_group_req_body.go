package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupReqBody 更新用户组请求体
type UpdateGroupReqBody struct {

	// 更新的用户组属性列表
	Operations []AttributeOperationDto `json:"operations"`
}

func (o UpdateGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupReqBody struct{}"
	}

	return strings.Join([]string{"UpdateGroupReqBody", string(data)}, " ")
}
