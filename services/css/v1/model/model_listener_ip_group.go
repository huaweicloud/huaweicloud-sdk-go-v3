package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListenerIpGroup struct {

	// **参数解释**： 访问控制组的类型。 **取值范围**： - white:白名单，只允许指定IP访问。 - black:黑名单，不允许指定IP访问。
	Type *string `json:"type,omitempty"`

	// 监听器关联的访问控制组的ID。创建时必选，更新时非必选。
	IpgroupId *string `json:"ipgroup_id,omitempty"`

	// 访问控制组的状态。
	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`
}

func (o ListenerIpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerIpGroup struct{}"
	}

	return strings.Join([]string{"ListenerIpGroup", string(data)}, " ")
}
