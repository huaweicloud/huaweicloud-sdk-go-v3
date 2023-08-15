package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBatchFailure struct {

	// 删除失败的ACL策略ID
	AclId *string `json:"acl_id,omitempty"`

	// 删除失败的ACL策略名称
	AclName *string `json:"acl_name,omitempty"`

	// 删除失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 删除失败的错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AclBatchFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBatchFailure struct{}"
	}

	return strings.Join([]string{"AclBatchFailure", string(data)}, " ")
}
