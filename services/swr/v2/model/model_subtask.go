package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Subtask struct {

	// 子任务ID
	Id int32 `json:"id"`

	// 内部任务ID
	JobId int32 `json:"job_id"`

	// 命名空间名
	Namespace string `json:"namespace"`

	// 制品仓库
	Repository string `json:"repository"`

	// 制品版本
	Tag string `json:"tag"`

	// sha256值
	Digest string `json:"digest"`

	// 老化动作， DEL
	Action string `json:"action"`

	// 状态，Initialized、Pending、InProgress、Succeed、Failed、Stopped
	Status string `json:"status"`

	// 状态信息
	StatusText string `json:"status_text"`

	// 开始时间
	OpTime string `json:"op_time"`
}

func (o Subtask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subtask struct{}"
	}

	return strings.Join([]string{"Subtask", string(data)}, " ")
}
