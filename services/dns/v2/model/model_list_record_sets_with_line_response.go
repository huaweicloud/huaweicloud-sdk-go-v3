package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordSetsWithLineResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Recordsets *[]QueryRecordSetWithLineAndTagsResp `json:"recordsets,omitempty" xml:"recordsets"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsWithLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineResponse", string(data)}, " ")
}
