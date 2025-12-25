package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSearchConditionRequestBody struct {

	// 检索条件名称
	ConditionName string `json:"condition_name"`

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 查询语句
	Query string `json:"query"`
}

func (o CreateSearchConditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchConditionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSearchConditionRequestBody", string(data)}, " ")
}
