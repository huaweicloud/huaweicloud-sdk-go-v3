package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindDbEncryptEipRequest Request Object
type BindDbEncryptEipRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	Body *BindEipRequest `json:"body,omitempty"`
}

func (o BindDbEncryptEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDbEncryptEipRequest struct{}"
	}

	return strings.Join([]string{"BindDbEncryptEipRequest", string(data)}, " ")
}
