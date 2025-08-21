package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryInheritSettingsBodyDto struct {

	// **参数解释：** 继承设置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Data *[]RepositoryInheritSettingUpdateBodyDto `json:"data,omitempty"`
}

func (o RepositoryInheritSettingsBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryInheritSettingsBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryInheritSettingsBodyDto", string(data)}, " ")
}
