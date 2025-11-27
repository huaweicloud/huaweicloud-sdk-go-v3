package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationViewQueryResponseData 描述
type ApplicationViewQueryResponseData struct {

	// **参数解释：** CMDB分配的uuid。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 应用或分组或组件的名称。 **取值范围：** 字符串，长度在[3,50]之间。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 应用或分组或组件code。 **取值范围：** 字符串，长度在[3,50]之间。
	Code *string `json:"code,omitempty"`

	// **参数解释：** 应用或分组或组件。 **取值范围：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 父节点id，即查询结果所在路径的父节点id。 **取值范围：** 字符串，长度24。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释：** 组件id。 **取值范围：** 字符串，长度24。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释：** 应用id。 **取值范围：** 字符串，长度24。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** 节点所在路径，由应用，组件，分组等id组成。 **取值范围：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 云厂商信息。 **取值范围：** - RMS：华为云厂商。 - ALI：阿里云厂商。 - OTHER：其他厂商。
	Vendor *string `json:"vendor,omitempty"`

	// **参数解释：** 跨账号资源所属的domainId。 **取值范围：** 不涉及。
	RelatedDomainId *string `json:"related_domain_id,omitempty"`
}

func (o ApplicationViewQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationViewQueryResponseData struct{}"
	}

	return strings.Join([]string{"ApplicationViewQueryResponseData", string(data)}, " ")
}
