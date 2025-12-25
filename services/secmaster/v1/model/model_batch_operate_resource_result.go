package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateResourceResult 批量操作资产返回对象
type BatchOperateResourceResult struct {

	// 失败id
	ErrorIds *[]string `json:"error_ids,omitempty"`

	// 成功id
	SuccessIds *[]string `json:"success_ids,omitempty"`
}

func (o BatchOperateResourceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateResourceResult struct{}"
	}

	return strings.Join([]string{"BatchOperateResourceResult", string(data)}, " ")
}
