package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStarrocksInstanceRequest Request Object
type CreateStarrocksInstanceRequest struct {

	// GaussDBForMySQL数据库实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *StarRocksCreateRequest `json:"body,omitempty"`
}

func (o CreateStarrocksInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStarrocksInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateStarrocksInstanceRequest", string(data)}, " ")
}
