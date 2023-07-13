package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIp
type PublicIp struct {

	// 弹性IP绑定类型，取值如下。 - auto_assign：自动绑定。 - bind_existing：使用已有。
	PublicBindType *string `json:"publicBindType,omitempty"`

	// 弹性IP的id，当publicBindType设置为bind_existing时，该值为用户某个已创建但尚未绑定的EIP的ID；当publicBindType设置为auto_assign时，该值设置为空。
	EipId *string `json:"eipId,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
