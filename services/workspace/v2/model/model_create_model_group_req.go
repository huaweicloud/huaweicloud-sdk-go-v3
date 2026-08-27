package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelGroupReq 创建模型组请求。
type CreateModelGroupReq struct {

	// 分组名称。
	Name string `json:"name"`

	// 分组描述。
	Description *string `json:"description,omitempty"`

	// 初始关联的供应商ID列表（可选）。
	ProviderIds *[]string `json:"provider_ids,omitempty"`
}

func (o CreateModelGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelGroupReq struct{}"
	}

	return strings.Join([]string{"CreateModelGroupReq", string(data)}, " ")
}
