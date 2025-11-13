package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePublicRecordsetsTaskRequestBody struct {

	// **参数解释：** 托管该记录的域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneNames []string `json:"zone_names"`

	Filter *BatchUpdatePublicRecordsetsFilter `json:"filter"`

	UpdateInfo *BatchUpdatePublicRecordsetsUpdateValue `json:"update_info"`
}

func (o BatchUpdatePublicRecordsetsTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePublicRecordsetsTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePublicRecordsetsTaskRequestBody", string(data)}, " ")
}
