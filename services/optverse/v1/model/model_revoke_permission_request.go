package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokePermissionRequest Request Object
type RevokePermissionRequest struct {

	// **参数解释**： 桶名。 **约束限制**： 不涉及 **取值范围**： 仅支持小写字母、数字、中划线和下划线英文句号，长度为[1-63]个字符。 **默认取值**： 不涉及
	Bucket string `json:"bucket"`
}

func (o RevokePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokePermissionRequest struct{}"
	}

	return strings.Join([]string{"RevokePermissionRequest", string(data)}, " ")
}
