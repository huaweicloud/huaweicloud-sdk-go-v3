package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRecordSetWithLineResponse Response Object
type BatchDeleteRecordSetWithLineResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 记录集列表信息。 **取值范围：** 不涉及。
	Recordsets *[]QueryRecordSetWithLineResp `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteRecordSetWithLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetWithLineResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetWithLineResponse", string(data)}, " ")
}
