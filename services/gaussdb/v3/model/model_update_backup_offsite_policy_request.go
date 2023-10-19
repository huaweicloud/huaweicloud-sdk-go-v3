package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupOffsitePolicyRequest Request Object
type UpdateBackupOffsitePolicyRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *UpdateBackupOffsitePolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateBackupOffsitePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupOffsitePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateBackupOffsitePolicyRequest", string(data)}, " ")
}
