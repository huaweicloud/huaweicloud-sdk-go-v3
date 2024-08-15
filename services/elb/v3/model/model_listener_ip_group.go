package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListenerIpGroup listener对象中的ipgroup信息
type ListenerIpGroup struct {

	// 参数解释：监听器关联的访问控制组的id。 创建时必选，更新时非必选。  约束限制： 指定的ipgroup必须已存在，不能指定为null，否则与enable_ipgroup冲突。
	IpgroupId string `json:"ipgroup_id"`

	// 参数解释：访问控制组的状态。 开启访问控制的监听器，允许直接删除。  取值范围： - true:开启访问控制。 - flase：关闭访问控制。
	EnableIpgroup bool `json:"enable_ipgroup"`

	// 参数解释：访问控制组的类型。  取值范围： - white:白名单，只允许指定ip访问。 - black:黑名单，不允许指定ip访问。
	Type string `json:"type"`
}

func (o ListenerIpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerIpGroup struct{}"
	}

	return strings.Join([]string{"ListenerIpGroup", string(data)}, " ")
}
