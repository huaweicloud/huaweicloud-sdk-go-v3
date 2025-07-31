package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordSetsResponse Response Object
type ListRecordSetsResponse struct {
	Links *PageLink `json:"links,omitempty"`

	// **参数解释：** 记录集列表信息。 **取值范围：** 不涉及。
	Recordsets *[]ListRecordSetsWithTags `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsResponse", string(data)}, " ")
}
