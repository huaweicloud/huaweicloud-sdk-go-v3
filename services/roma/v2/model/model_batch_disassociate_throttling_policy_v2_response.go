package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDisassociateThrottlingPolicyV2Response struct {

	// 成功解除绑定的API和流控策略绑定关系的数量
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 解除绑定失败的API和流控绑定关系及错误信息
	Failure        *[]ThrottleBindingBatchFailure `json:"failure,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o BatchDisassociateThrottlingPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateThrottlingPolicyV2Response struct{}"
	}

	return strings.Join([]string{"BatchDisassociateThrottlingPolicyV2Response", string(data)}, " ")
}
