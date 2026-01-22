package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddBlackWhiteListBatchVo struct {

	// **参数解释**： 黑白名单重复添加列表 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	DuplicatedList *[]BlackWhiteInfo `json:"duplicated_list,omitempty"`

	// **参数解释**： 黑白名单添加失败列表 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	FailedList *[]BlackWhiteInfo `json:"failed_list,omitempty"`

	// **参数解释**： 黑白名单添加成功列表 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	SuccessList *[]BlackWhiteInfo `json:"success_list,omitempty"`
}

func (o AddBlackWhiteListBatchVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteListBatchVo struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteListBatchVo", string(data)}, " ")
}
