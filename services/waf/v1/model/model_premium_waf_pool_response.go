package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PremiumWafPoolResponse struct {

	// **参数解释：** 实例组ID，用于唯一标识一个实例组。 **取值范围：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 实例组名称。 **取值范围：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 实例组所在的区域（Region）。 **取值范围：** 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释：** 实例组类型。 **取值范围：** - elb：基础elb类型 - elb-v2：elb-v2类型 - elb-shadow：saas化elb类型 - standard-container：反向代理独享引擎组（云内，承载租户专用） - standard-cloud：反向代理独享引擎组（云内） - standard：反向代理独享引擎组（云外） - detector-cloud：旁路检测独享引擎组（云内） - detector：旁路检测独享引擎组（云外）
	Type *string `json:"type,omitempty"`

	// **参数解释：** 实例组关联的虚拟私有云ID。 **取值范围：** 不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 实例组的描述信息。 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 实例组关联的防护域名列表。 **取值范围：** 不涉及
	Hosts *[]IdNameEntry `json:"hosts,omitempty"`

	// **参数解释：** 实例组关联的引擎实例列表。 **取值范围：** 不涉及
	Instances *[]IdNameEntry `json:"instances,omitempty"`

	// **参数解释：** 实例组关联的企业项目ID。 **取值范围：** - 0：代表default企业项目 - 其他为企业项目ID，长度为36个字符
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 实例组创建时间。Unix时间戳格式，单位为毫秒（ms）。 **取值范围：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o PremiumWafPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafPoolResponse struct{}"
	}

	return strings.Join([]string{"PremiumWafPoolResponse", string(data)}, " ")
}
