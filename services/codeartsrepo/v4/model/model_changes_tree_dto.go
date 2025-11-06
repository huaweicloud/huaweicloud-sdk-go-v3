package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangesTreeDto 合并请求文件变更列表详情
type ChangesTreeDto struct {

	// **参数解释：** 展示审核文件。
	CanShowMyApprovalFiles *bool `json:"can_show_my_approval_files,omitempty"`

	// **参数解释：** 变更树。
	Tree *[]ChangesTreeObjectDto `json:"tree,omitempty"`
}

func (o ChangesTreeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangesTreeDto struct{}"
	}

	return strings.Join([]string{"ChangesTreeDto", string(data)}, " ")
}
