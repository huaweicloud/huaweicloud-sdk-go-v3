package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelBatchCreateReq 批量新增模型请求。
type ModelBatchCreateReq struct {

	// 模型列表。
	Models []CreateModelReq `json:"models"`
}

func (o ModelBatchCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelBatchCreateReq struct{}"
	}

	return strings.Join([]string{"ModelBatchCreateReq", string(data)}, " ")
}
