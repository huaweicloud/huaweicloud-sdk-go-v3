package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartStarrocksInstanceRequest Request Object
type RestartStarrocksInstanceRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	StarrocksInstanceId string `json:"starrocks_instance_id"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o RestartStarrocksInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartStarrocksInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartStarrocksInstanceRequest", string(data)}, " ")
}
