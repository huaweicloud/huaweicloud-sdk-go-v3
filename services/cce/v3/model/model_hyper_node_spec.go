package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperNodeSpec struct {

	// **参数解释**： 超节点规格
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**： 所属节点池ID
	NodepoolID *string `json:"nodepoolID,omitempty"`

	// **参数解释**： 超节点下节点相关的配置。
	NodeTemplate *[]NodeTemplateInHyperNode `json:"nodeTemplate,omitempty"`

	// **参数解释**： 付费方式 **取值范围**： - prepaid: 预付费，即包年包月； - postpaid: 后付费，即按需付费；
	ChargeMode *string `json:"chargeMode,omitempty"`
}

func (o HyperNodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeSpec struct{}"
	}

	return strings.Join([]string{"HyperNodeSpec", string(data)}, " ")
}
