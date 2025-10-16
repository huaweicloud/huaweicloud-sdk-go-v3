package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordInfo struct {

	// **参数解释：** 解析记录的主机记录。 **取值范围：** 不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释：** 解析记录的记录值。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o RecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfo struct{}"
	}

	return strings.Join([]string{"RecordInfo", string(data)}, " ")
}
