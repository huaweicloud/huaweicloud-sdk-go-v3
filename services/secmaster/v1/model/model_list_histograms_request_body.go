package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListHistogramsRequestBody struct {

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	// 查询开始时间点
	From int64 `json:"from"`

	// 数据管道ID
	PipeId string `json:"pipe_id"`

	// 查询语句
	Query string `json:"query"`

	// 查询结束时间点
	To int64 `json:"to"`
}

func (o ListHistogramsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistogramsRequestBody struct{}"
	}

	return strings.Join([]string{"ListHistogramsRequestBody", string(data)}, " ")
}
