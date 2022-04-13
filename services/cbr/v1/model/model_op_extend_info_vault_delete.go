package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoVaultDelete struct {
	// 删除失败的资源数量

	FailCount *int32 `json:"fail_count,omitempty"`
	// 删除的备份数量

	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o OpExtendInfoVaultDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoVaultDelete struct{}"
	}

	return strings.Join([]string{"OpExtendInfoVaultDelete", string(data)}, " ")
}
