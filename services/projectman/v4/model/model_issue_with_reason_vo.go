package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueWithReasonVo struct {

	// **参数解释**： 工作项唯一ID。可以通过[查询工作项列表](ListIpdProjectIssues.xml)或者[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**id**字段的值就是工作项ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作项类型。 **取值范围**： 支持多种工作项类型，使用英文逗号分隔。 - 系统设备类项目：RR、SF、IR、SR、AR、Task、Bug - 独立软件类项目：RR、SF、IR、US、Task、Bug - 云服务类项目：RR、Epic、FE、US、Task、Bug
	Category *string `json:"category,omitempty"`

	// **参数解释**： 工作项标题。  **取值范围**： 不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释**： 工作项唯一编码number。  **取值范围**： 不涉及。
	Number *string `json:"number,omitempty"`

	// **参数解释**： 操作失败原因。  **取值范围**： 不涉及。
	Reason *string `json:"reason,omitempty"`
}

func (o IssueWithReasonVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueWithReasonVo struct{}"
	}

	return strings.Join([]string{"IssueWithReasonVo", string(data)}, " ")
}
