package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Severity 工作项重要程度
type Severity struct {

	// **参数解释：** 工作项的重要程度。 **取值范围：** - 关键 - 重要 - 一般 - 提示
	Name *string `json:"name,omitempty"`

	// **参数解释：** 重要程度id。 **取值范围：** 10 （关键） 11 （重要） 12 （一般） 13 （提示）
	Id *int32 `json:"id,omitempty"`
}

func (o Severity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Severity struct{}"
	}

	return strings.Join([]string{"Severity", string(data)}, " ")
}
