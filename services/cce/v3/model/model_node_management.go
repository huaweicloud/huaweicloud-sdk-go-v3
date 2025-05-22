package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeManagement struct {

	// **参数解释**： 云服务器组ID，指定后将绑定云服务器组，节点池中所有节点将创建在该云服务器组下。绑定云服务器组后节点池中的节点数量不允许超出云服务器组可添加的云服务器数量，否则将导致节点池无法扩容。 > - 绑定云服务器组后，云服务器将严格按照亲和策略分布，同时也会限制节点池中节点个数上限。由于ECS创建云服务器时本身具有一定反亲和能力，如果仅需云服务器分散的创建在不同主机上以提高业务的可靠性，又不希望节点个数受到云服务器组的限制，请勿绑定云服务器组。 > - 云服务组支持解绑，解绑后存量节点仍属于原云服务器组，新建节点将不再绑定云服务器组。 当节点池中不存在节点时支持绑定新的云服务器组或者切换绑定的云服务器组  **约束限制**： 指定云服务器组时节点池中的节点数量不允许超出云服务器组的配额限制。 **取值范围**： - 不指定或者指定空字符串：表示解绑云服务器组 - 云服务器组ID：表示切换节点池绑定的云服务组  **默认取值**： 不涉及
	ServerGroupReference *string `json:"serverGroupReference,omitempty"`
}

func (o NodeManagement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeManagement struct{}"
	}

	return strings.Join([]string{"NodeManagement", string(data)}, " ")
}
