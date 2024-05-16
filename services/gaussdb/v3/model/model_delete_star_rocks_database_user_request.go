package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRocksDatabaseUserRequest Request Object
type DeleteStarRocksDatabaseUserRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库账户名。
	UserName string `json:"user_name"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o DeleteStarRocksDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRocksDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteStarRocksDatabaseUserRequest", string(data)}, " ")
}
