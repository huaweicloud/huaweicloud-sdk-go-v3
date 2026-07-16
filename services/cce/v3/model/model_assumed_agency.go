package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumedAgency 委托元数据信息
type AssumedAgency struct {

	// **参数解释**： 委托的唯一身份标识信息，形如： sts::{account_id}::assumed-agency:{agency_name}/{agency_session_name} **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Urn *string `json:"urn,omitempty"`

	// **参数解释：** 委托的id属性，形如： {agency_id}:{agency_session_name} **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`
}

func (o AssumedAgency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumedAgency struct{}"
	}

	return strings.Join([]string{"AssumedAgency", string(data)}, " ")
}
