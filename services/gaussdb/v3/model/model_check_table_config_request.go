package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableConfigRequest Request Object
type CheckTableConfigRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *TableConfigCheckRequestV3 `json:"body,omitempty"`
}

func (o CheckTableConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableConfigRequest struct{}"
	}

	return strings.Join([]string{"CheckTableConfigRequest", string(data)}, " ")
}
