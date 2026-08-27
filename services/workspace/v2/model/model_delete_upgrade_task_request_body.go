package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUpgradeTaskRequestBody 批量删除升级任务请求
type DeleteUpgradeTaskRequestBody struct {

	// 任务ID列表
	TaskIds *[]string `json:"task_ids,omitempty"`
}

func (o DeleteUpgradeTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUpgradeTaskRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteUpgradeTaskRequestBody", string(data)}, " ")
}
