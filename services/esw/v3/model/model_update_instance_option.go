package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceOption struct {

	// - 参数解释：ESW实例描述信息。 - 约束限制：   - 长度范围为0~255个字符。   - 不能包含“<”和“>”。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Description *string `json:"description,omitempty"`

	// - 参数解释：ESW实例名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o UpdateInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceOption struct{}"
	}

	return strings.Join([]string{"UpdateInstanceOption", string(data)}, " ")
}
