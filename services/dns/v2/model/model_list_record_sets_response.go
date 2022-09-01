package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordSetsResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Recordsets *[]ListRecordSetsWithTags `json:"recordsets,omitempty" xml:"recordsets"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsResponse", string(data)}, " ")
}
