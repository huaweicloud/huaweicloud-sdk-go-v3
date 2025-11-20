package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Premium struct {

	// **参数解释：** 是否开通独享模式 **取值范围：** 不涉及
	Purchased *bool `json:"purchased,omitempty"`

	// **参数解释：** 独享实例数量，包括elb **取值范围：** 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** elb实例数量 **取值范围：** 不涉及
	Elb *int32 `json:"elb,omitempty"`

	// **参数解释：** 独享WAF实例数量 **取值范围：** 不涉及
	Dedicated *int32 `json:"dedicated,omitempty"`
}

func (o Premium) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Premium struct{}"
	}

	return strings.Join([]string{"Premium", string(data)}, " ")
}
