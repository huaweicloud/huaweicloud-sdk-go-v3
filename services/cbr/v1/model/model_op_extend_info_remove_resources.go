package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoRemoveResources struct {

	// 删除失败的资源数量
	FailCount *int32 `json:"fail_count,omitempty" xml:"fail_count"`

	// 删除的备份数量
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	//
	Resources *[]Resource `json:"resources,omitempty" xml:"resources"`
}

func (o OpExtendInfoRemoveResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoRemoveResources struct{}"
	}

	return strings.Join([]string{"OpExtendInfoRemoveResources", string(data)}, " ")
}
