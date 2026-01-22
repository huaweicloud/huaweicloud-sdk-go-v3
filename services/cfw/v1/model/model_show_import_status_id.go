package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowImportStatusId struct {

	// **参数解释**： 防护对象id **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规则导入状态 **约束限制**： 不涉及 **取值范围**： 0: 无任务态 1: 任务等待 2: 任务执行 3: 任务成功 4: 任务失败 5: 任务部分成功 6: 任务全部失败 **默认取值**： 不涉及
	Status *int32 `json:"status,omitempty"`
}

func (o ShowImportStatusId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImportStatusId struct{}"
	}

	return strings.Join([]string{"ShowImportStatusId", string(data)}, " ")
}
