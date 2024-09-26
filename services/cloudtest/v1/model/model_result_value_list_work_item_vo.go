package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueListWorkItemVo 请求的返回的数据对象
type ResultValueListWorkItemVo struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value *[]WorkItemVo `json:"value,omitempty"`
}

func (o ResultValueListWorkItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListWorkItemVo struct{}"
	}

	return strings.Join([]string{"ResultValueListWorkItemVo", string(data)}, " ")
}
