package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AffinityResponse **参数解释：** 节点亲和类型。
type AffinityResponse struct {

	// **参数解释：** 节点亲和类型。 **取值范围：** - AFFINITY：亲和。 - ANTI_AFFINITY：反亲和。
	AffinityType string `json:"affinity_type"`

	// **参数解释：** 是否设置强亲和。 **取值范围：** - true：设置强亲和。 - false：不设置强亲和。
	Required bool `json:"required"`

	// **参数解释：** 选择节点方式。 **取值范围：** IP。
	SelectionMode string `json:"selection_mode"`

	// **参数解释：** 通过上述方式选择的列表，长度不能超过20。
	Targets map[string]string `json:"targets"`
}

func (o AffinityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AffinityResponse struct{}"
	}

	return strings.Join([]string{"AffinityResponse", string(data)}, " ")
}
