package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryParentDto struct {

	// 子节点实例ID。
	ChildId string `json:"childId"`
}

func (o QueryParentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryParentDto struct{}"
	}

	return strings.Join([]string{"QueryParentDto", string(data)}, " ")
}
