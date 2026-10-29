package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProgressDetailV2 struct {

	// **参数解释:** 检查进度百分比。 **取值范围:** 检查中、检查成功时，为0%-100%； 检查失败、任务中止、未检查时，为null。
	Ratio *string `json:"ratio,omitempty"`

	// **参数解释:** 检查执行状态。 **取值范围:** - running：检查中。 - success：检查成功。 - null：检查失败、任务中止、未检查。
	Info *string `json:"info,omitempty"`
}

func (o ProgressDetailV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgressDetailV2 struct{}"
	}

	return strings.Join([]string{"ProgressDetailV2", string(data)}, " ")
}
