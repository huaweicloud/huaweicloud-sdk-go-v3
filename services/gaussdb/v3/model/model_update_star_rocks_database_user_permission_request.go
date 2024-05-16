package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStarRocksDatabaseUserPermissionRequest Request Object
type UpdateStarRocksDatabaseUserPermissionRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *StarRocksDatabaseUserPSinfo `json:"body,omitempty"`
}

func (o UpdateStarRocksDatabaseUserPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarRocksDatabaseUserPermissionRequest struct{}"
	}

	return strings.Join([]string{"UpdateStarRocksDatabaseUserPermissionRequest", string(data)}, " ")
}
