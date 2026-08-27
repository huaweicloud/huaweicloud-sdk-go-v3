package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveModelConfigReq 批量移除模型分组与资源关联请求。
type RemoveModelConfigReq struct {

	// 关联项列表，每项指定要移除的模型分组与资源的关联。
	Items []ModelConfigItem `json:"items"`
}

func (o RemoveModelConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveModelConfigReq struct{}"
	}

	return strings.Join([]string{"RemoveModelConfigReq", string(data)}, " ")
}
