package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsMatch Tms通过标签查询资源时，传入的特殊key-value匹配项，目前只支持传入资源名称进行模糊查询。
type TmsMatch struct {

	// **参数解释：** 匹配项名称，目前只支持resource_name。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 匹配项的值，不区分大小写，key为resource_name时使用模糊查询，匹配资源名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Value string `json:"value"`
}

func (o TmsMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsMatch struct{}"
	}

	return strings.Join([]string{"TmsMatch", string(data)}, " ")
}
