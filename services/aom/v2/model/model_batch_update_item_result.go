package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateItemResult struct {

	// 资源名称别名。
	ResourceAlias *string `json:"resource_alias,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源更新状态。
	RetStatus *string `json:"ret_status,omitempty"`
}

func (o BatchUpdateItemResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateItemResult struct{}"
	}

	return strings.Join([]string{"BatchUpdateItemResult", string(data)}, " ")
}
