package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchConditionResponse Response Object
type ShowSearchConditionResponse struct {

	// 检索条件ID
	ConditionId *string `json:"condition_id,omitempty"`

	// 检索条件名称
	ConditionName *string `json:"condition_name,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 查询语句
	Query          *string `json:"query,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSearchConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchConditionResponse struct{}"
	}

	return strings.Join([]string{"ShowSearchConditionResponse", string(data)}, " ")
}
