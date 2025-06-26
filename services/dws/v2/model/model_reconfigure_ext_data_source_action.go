package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReconfigureExtDataSourceAction **参数解释**： 更新数据源配置。 **取值范围**： 不涉及。
type ReconfigureExtDataSourceAction struct {

	// **参数解释**： 重启。 **取值范围**： 不涉及。
	Reboot *bool `json:"reboot,omitempty"`

	// **参数解释**： 委托。 **取值范围**： 不涉及。
	Agency *string `json:"agency,omitempty"`
}

func (o ReconfigureExtDataSourceAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReconfigureExtDataSourceAction struct{}"
	}

	return strings.Join([]string{"ReconfigureExtDataSourceAction", string(data)}, " ")
}
