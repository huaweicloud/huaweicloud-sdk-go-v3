package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAppToGroupRequestBody 移动应用至指定分组请求体
type MoveAppToGroupRequestBody struct {

	// 分组id
	GroupId string `json:"group_id"`

	// 应用id列表
	ApplicationIds []string `json:"application_ids"`
}

func (o MoveAppToGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppToGroupRequestBody struct{}"
	}

	return strings.Join([]string{"MoveAppToGroupRequestBody", string(data)}, " ")
}
