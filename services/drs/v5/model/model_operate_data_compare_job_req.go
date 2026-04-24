package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateDataCompareJobReq 操作对比任务请求体
type OperateDataCompareJobReq struct {

	// 操作的对比任务ID列表。
	CompareJobIds []string `json:"compare_job_ids"`
}

func (o OperateDataCompareJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateDataCompareJobReq struct{}"
	}

	return strings.Join([]string{"OperateDataCompareJobReq", string(data)}, " ")
}
