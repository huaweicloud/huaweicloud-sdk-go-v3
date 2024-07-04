package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStarrocksParamsRequest Request Object
type UpdateStarrocksParamsRequest struct {

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *UpdateParamInfo `json:"body,omitempty"`
}

func (o UpdateStarrocksParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarrocksParamsRequest struct{}"
	}

	return strings.Join([]string{"UpdateStarrocksParamsRequest", string(data)}, " ")
}
