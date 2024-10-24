package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserReqBody 更新用户请求体
type UpdateUserReqBody struct {

	// 更新的用户属性列表
	Operations []AttributeOperationDto `json:"operations"`
}

func (o UpdateUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReqBody struct{}"
	}

	return strings.Join([]string{"UpdateUserReqBody", string(data)}, " ")
}
