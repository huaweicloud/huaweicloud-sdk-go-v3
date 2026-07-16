package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateCertNodeList 轮转节点证书参数。集群内已有节点轮转证书失败，可以通过此接口重试。
type RotateCertNodeList struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： List
	Kind string `json:"kind"`

	// **参数解释**： 需轮转证书的节点列表，当前最多支持同时轮转200个节点。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeList []RotateCertNode `json:"nodeList"`
}

func (o RotateCertNodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateCertNodeList struct{}"
	}

	return strings.Join([]string{"RotateCertNodeList", string(data)}, " ")
}
