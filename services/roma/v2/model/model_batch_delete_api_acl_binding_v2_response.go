package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteApiAclBindingV2Response struct {

	// 成功解除绑定的ACL策略数量
	SuccessCount *int32 `json:"success_count,omitempty" xml:"success_count"`

	// 解除绑定失败的ACL策略及错误信息
	Failure        *[]AclBindingBatchFailure `json:"failure,omitempty" xml:"failure"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchDeleteApiAclBindingV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteApiAclBindingV2Response struct{}"
	}

	return strings.Join([]string{"BatchDeleteApiAclBindingV2Response", string(data)}, " ")
}
