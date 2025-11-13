package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEmailRecordSetResponse Response Object
type ShowEmailRecordSetResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 记录集列表信息。 **取值范围：** 不涉及。
	Recordsets *[]ListRecordSets `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowEmailRecordSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEmailRecordSetResponse struct{}"
	}

	return strings.Join([]string{"ShowEmailRecordSetResponse", string(data)}, " ")
}
