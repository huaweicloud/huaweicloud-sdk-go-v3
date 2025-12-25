package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateBaselineResult 失败或者成功的id
type BatchOperateBaselineResult struct {

	// 成功id列表
	ErrorIds []string `json:"error_ids"`

	// 失败id列表
	SuccessIds []string `json:"success_ids"`
}

func (o BatchOperateBaselineResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateBaselineResult struct{}"
	}

	return strings.Join([]string{"BatchOperateBaselineResult", string(data)}, " ")
}
