package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateJobRsp struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	Status *BatchOperateJobStatus `json:"status,omitempty"`

	// 操作结果失败信息，仅在操作失败时会返回
	Message *string `json:"message,omitempty"`
}

func (o BatchOperateJobRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateJobRsp struct{}"
	}

	return strings.Join([]string{"BatchOperateJobRsp", string(data)}, " ")
}
