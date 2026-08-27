package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelBatchDeleteReq 批量删除模型请求。
type ModelBatchDeleteReq struct {

	// 模型id列表。
	ModelIds []string `json:"model_ids"`
}

func (o ModelBatchDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelBatchDeleteReq struct{}"
	}

	return strings.Join([]string{"ModelBatchDeleteReq", string(data)}, " ")
}
