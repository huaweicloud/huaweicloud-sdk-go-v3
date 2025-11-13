package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeletePublicRecordsetsTaskRequestBody struct {

	// **参数解释：** 托管该记录的域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneNames []string `json:"zone_names"`

	// **参数解释：** 主机记录。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RecordsetNamePrefixes []string `json:"recordset_name_prefixes"`
}

func (o BatchDeletePublicRecordsetsTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicRecordsetsTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicRecordsetsTaskRequestBody", string(data)}, " ")
}
