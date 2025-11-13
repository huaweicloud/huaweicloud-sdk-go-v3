package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowBatchOperationTaskErrorItem struct {

	// **参数解释：** 错误记录。 **取值范围：** 不涉及。
	Item *string `json:"item,omitempty"`

	// **参数解释：** 错误信息。 **取值范围：** 不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ShowBatchOperationTaskErrorItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchOperationTaskErrorItem struct{}"
	}

	return strings.Join([]string{"ShowBatchOperationTaskErrorItem", string(data)}, " ")
}
