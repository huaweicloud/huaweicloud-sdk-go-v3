package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListenerIpGroup struct {

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
