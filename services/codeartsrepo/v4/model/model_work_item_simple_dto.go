package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemSimpleDto **参数解释：** 关联工作项基本信息。
type WorkItemSimpleDto struct {

	// **参数解释：** 工作项编号。
	RelatedId *string `json:"related_id,omitempty"`

	// **参数解释：** 工作项地址。
	RelatedUrl *string `json:"related_url,omitempty"`
}

func (o WorkItemSimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemSimpleDto struct{}"
	}

	return strings.Join([]string{"WorkItemSimpleDto", string(data)}, " ")
}
