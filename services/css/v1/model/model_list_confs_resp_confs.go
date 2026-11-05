package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConfsRespConfs struct {

	// **参数解释**： 配置文件名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 配置文件状态。 **取值范围**： - checking：正在检查中。 - available：可用。 - unavailable：不可用。 - starting：正开始。 - working：工作中。 - stopping：正在停止。 - stopped：已经停止。 - restoring：正在回滚。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 配置文件内容。 **取值范围**： 不涉及
	ConfContent *string `json:"confContent,omitempty"`

	Setting *ListConfsRespConfsSetting `json:"setting,omitempty"`

	// **参数解释**： 更新时间，格式：yyyy-MM-dd'T'HH:mm:ss。 **取值范围**： 不涉及
	UpdateAt *string `json:"updateAt,omitempty"`

	// **参数解释**： 配置文件描述。 **取值范围**： 长度不超过128个字符。
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 是否隐藏敏感词。 **取值范围**： - true：隐藏敏感词。 - false：未隐藏敏感词。
	SensitiveWordsShielding *bool `json:"sensitiveWordsShielding,omitempty"`
}

func (o ListConfsRespConfs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfsRespConfs struct{}"
	}

	return strings.Join([]string{"ListConfsRespConfs", string(data)}, " ")
}
