package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportDlqMessageReq struct {

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 消息ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	MsgIdList []string `json:"msg_id_list"`

	// **参数解释**： 唯一Key列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	UniqKeyList *[]string `json:"uniq_key_list,omitempty"`
}

func (o ExportDlqMessageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageReq struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageReq", string(data)}, " ")
}
