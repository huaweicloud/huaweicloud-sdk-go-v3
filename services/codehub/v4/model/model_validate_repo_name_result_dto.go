package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidateRepoNameResultDto struct {

	// **参数解释：** 仓库名。 **约束限制：** - 必填。 - 大小写字母、数字、下划线开头，可包含大小写字母、数字、中划线、下划线、英文句点，但不能以.git、.atom或.结尾 - 代码总路径长度（代码组名称和仓库名称总长度）不超过256字符 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 项目ID。 **约束限制：** 必填。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 代码组ID，若需要检查的仓库名称在项目根目录下可不传此参数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 1-2147483647
	GroupId *int32 `json:"group_id,omitempty"`

	// **参数解释：** 是否校验通过 **取值范围：** - true，校验通过。 - false，校验未通过。 **约束限制：** 不涉及。
	Result *bool `json:"result,omitempty"`

	// **参数解释：** 异常信息 **约束限制：** 不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ValidateRepoNameResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRepoNameResultDto struct{}"
	}

	return strings.Join([]string{"ValidateRepoNameResultDto", string(data)}, " ")
}
