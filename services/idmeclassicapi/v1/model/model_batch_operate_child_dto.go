package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateChildDto struct {

	// 子节点实例ID列表。
	ChildList []string `json:"childList"`

	// 父节点实例ID。
	ParentId string `json:"parentId"`
}

func (o BatchOperateChildDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateChildDto struct{}"
	}

	return strings.Join([]string{"BatchOperateChildDto", string(data)}, " ")
}
