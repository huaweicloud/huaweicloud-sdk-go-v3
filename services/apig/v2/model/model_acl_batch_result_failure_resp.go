package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBatchResultFailureResp struct {

	// 删除失败的ACL策略ID
	AclId *string `json:"acl_id,omitempty"`

	// 删除失败的ACL策略名称
	AclName *string `json:"acl_name,omitempty"`

	// 删除失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 删除失败的错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AclBatchResultFailureResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBatchResultFailureResp struct{}"
	}

	return strings.Join([]string{"AclBatchResultFailureResp", string(data)}, " ")
}
