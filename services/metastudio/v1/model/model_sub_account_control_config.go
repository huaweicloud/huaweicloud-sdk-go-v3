package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubAccountControlConfig 子账号控制配置
type SubAccountControlConfig struct {

	// 子账号业务是否单独控制。
	SeparatelyControlled *bool `json:"separately_controlled,omitempty"`

	// 子账号类型。 * IAM_USER: 使用调用者IAM user id填充 X-App-UserId。如果请求中携带X-App-UserId，也会被忽略替换。
	SubAccountType *string `json:"sub_account_type,omitempty"`
}

func (o SubAccountControlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubAccountControlConfig struct{}"
	}

	return strings.Join([]string{"SubAccountControlConfig", string(data)}, " ")
}
