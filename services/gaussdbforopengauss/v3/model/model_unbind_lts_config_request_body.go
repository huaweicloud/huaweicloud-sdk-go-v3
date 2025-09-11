package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnbindLtsConfigRequestBody struct {

	// **参数解释**: 解除关联的实例日志列表。 **约束限制**: 不涉及。
	List []InstanceDeleteLtsConfigOption `json:"list"`
}

func (o UnbindLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UnbindLtsConfigRequestBody", string(data)}, " ")
}
