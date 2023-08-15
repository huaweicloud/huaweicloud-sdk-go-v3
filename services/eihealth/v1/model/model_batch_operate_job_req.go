package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchOperateJobReq struct {

	// 批量操作作业id列表
	Ids []string `json:"ids"`
}

func (o BatchOperateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateJobReq struct{}"
	}

	return strings.Join([]string{"BatchOperateJobReq", string(data)}, " ")
}
