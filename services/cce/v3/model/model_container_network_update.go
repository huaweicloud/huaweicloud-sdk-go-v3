package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerNetworkUpdate struct {

	// 容器网络网段列表。1.21及新版本集群，当集群网络类型为vpc-router时，支持增量添加容器网段，最多配置20个。  此参数在集群更新后不可更改，请谨慎选择。
	Cidrs *[]ContainerCidr `json:"cidrs,omitempty"`

	// **参数解释：** 容器网络网段列表。当集群网络类型为vpc-router时，支持添加或删除容器网段。支持的集群版本如下： - v1.28.15-r60及以上版本 - v1.29.15-r20及以上版本 - v1.30.14-r20及以上版本 - v1.31.10-r20及以上版本 - v1.32.6-r20及以上版本 - v1.33.5-r10及以上版本  支持修改集群容器网段的顺序，顺序在前的容器网段优先被使用。 **约束限制：** - 最多支持配置20个子网。 - 填写为空时，该字段不生效。  **取值范围：** 不涉及 **默认取值：** 不涉及
	Containercidrs *[]string `json:"containercidrs,omitempty"`
}

func (o ContainerNetworkUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNetworkUpdate struct{}"
	}

	return strings.Join([]string{"ContainerNetworkUpdate", string(data)}, " ")
}
