package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// listener对象中的ipgroup信息
type ListenerIpGroup struct {
	// 监听器关联的访问控制组的id。 创建时必选，更新时非必选。 指定的ipgroup必须已存在，不能指定为null，否则与enable_ipgroup冲突。

	IpgroupId string `json:"ipgroup_id"`
	// 访问控制组的状态。 True:开启访问控制； False：关闭访问控制； 开启访问控制的监听器，允许直接删除。

	EnableIpgroup bool `json:"enable_ipgroup"`
	// 访问控制组的类型。 white:白名单，只允许指定ip访问； black:黑名单，不允许指定ip访问；

	Type string `json:"type"`
}

func (o ListenerIpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerIpGroup struct{}"
	}

	return strings.Join([]string{"ListenerIpGroup", string(data)}, " ")
}
