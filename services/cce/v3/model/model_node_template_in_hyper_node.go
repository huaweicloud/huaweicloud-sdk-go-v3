package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTemplateInHyperNode 超节点下节点的信息。
type NodeTemplateInHyperNode struct {

	// **参数解释**： 超节点下节点所在的可用区。 [CCE支持的可用区请参考[地区和终端节点](https://console.huaweicloud.com/apiexplorer/#/endpoint/CCE)。](tag:hws) [CCE支持的可用区请参考[地区和终端节点](https://console-intl.huaweicloud.com/apiexplorer/#/endpoint/CCE)。](tag:hws_hk)
	Az *string `json:"az,omitempty"`

	// **参数解释**： 超节点下节点的操作系统类型。具体支持的操作系统请参见[节点操作系统说明](node-os.xml)。
	Os *string `json:"os,omitempty"`

	Login *Login `json:"login,omitempty"`

	RootVolume *Volume `json:"rootVolume,omitempty"`

	// **参数解释**： 节点的数据盘参数。
	DataVolumes *[]Volume `json:"dataVolumes,omitempty"`

	Storage *Storage `json:"storage,omitempty"`

	// **参数解释**： 超节点创建时下发到节点上的 k8s 标签，格式为key/value键值对。此接口中仅为展示作用。 示例： ``` \"k8sTags\": {   \"key\": \"value\" } ```
	K8sTags map[string]string `json:"k8sTags,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`

	ExtendParam *NodeExtendParam `json:"extendParam,omitempty"`

	HostnameConfig *HostnameConfig `json:"hostnameConfig,omitempty"`
}

func (o NodeTemplateInHyperNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTemplateInHyperNode struct{}"
	}

	return strings.Join([]string{"NodeTemplateInHyperNode", string(data)}, " ")
}
