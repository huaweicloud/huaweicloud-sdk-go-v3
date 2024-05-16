package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStarRocksDatabaseUserRequest Request Object
type CreateStarRocksDatabaseUserRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *StarRocksDatabaseUserInfo `json:"body,omitempty"`
}

func (o CreateStarRocksDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStarRocksDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"CreateStarRocksDatabaseUserRequest", string(data)}, " ")
}
