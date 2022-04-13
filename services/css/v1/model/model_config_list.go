package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigList struct {
	// 操作ID

	Id *string `json:"id,omitempty"`
	// 集群ID。

	ClusterId *string `json:"clusterId,omitempty"`
	// 创建时间。

	CreateAt *interface{} `json:"createAt,omitempty"`
	// 状态。

	Status *string `json:"status,omitempty"`
	// 结束时间。

	FinishedAt *interface{} `json:"finishedAt,omitempty"`
	// 修改参数配置记录。

	ModifyDeleteReset *string `json:"modifyDeleteReset,omitempty"`
	// 返回错误信息。

	FailedMsg *string `json:"failedMsg,omitempty"`
}

func (o ConfigList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigList struct{}"
	}

	return strings.Join([]string{"ConfigList", string(data)}, " ")
}
