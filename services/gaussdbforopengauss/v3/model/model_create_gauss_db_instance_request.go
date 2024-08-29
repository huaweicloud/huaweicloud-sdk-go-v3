package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussDbInstanceRequest Request Object
type CreateGaussDbInstanceRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 委托urn。使用RAM共享的KMS秘钥创建包周期实例时必填,格式iam::{account_id}:agency:{agency_name}。
	SubscriptionAgency *string `json:"Subscription-Agency,omitempty"`

	Body *OpenGaussInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateGaussDbInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussDbInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussDbInstanceRequest", string(data)}, " ")
}
