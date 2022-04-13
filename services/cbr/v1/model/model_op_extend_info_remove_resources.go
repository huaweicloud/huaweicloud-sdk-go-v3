package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoRemoveResources struct {
	// 删除失败的资源数量

	FailCount *int32 `json:"fail_count,omitempty"`
	// 删除的备份数量

	TotalCount *int32 `json:"total_count,omitempty"`
	//

	Resources *[]Resource `json:"resources,omitempty"`
}

func (o OpExtendInfoRemoveResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoRemoveResources struct{}"
	}

	return strings.Join([]string{"OpExtendInfoRemoveResources", string(data)}, " ")
}
