package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataBaseConfigRequest Request Object
type CheckDataBaseConfigRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *DbConfigCheckRequestV3 `json:"body,omitempty"`
}

func (o CheckDataBaseConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataBaseConfigRequest struct{}"
	}

	return strings.Join([]string{"CheckDataBaseConfigRequest", string(data)}, " ")
}
