package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopIpRequest Request Object
type ListTopIpRequest struct {

	// **参数解释：** 企业项目id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 起始时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From int64 `json:"from"`

	// **参数解释：** 结束时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To int64 `json:"to"`

	// **参数解释：** 要查询的前几的结果 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Top *int32 `json:"top,omitempty"`

	// **参数解释：** 要查询域名列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *string `json:"hosts,omitempty"`

	// **参数解释：** 要查询实例列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instances *string `json:"instances,omitempty"`
}

func (o ListTopIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopIpRequest struct{}"
	}

	return strings.Join([]string{"ListTopIpRequest", string(data)}, " ")
}
