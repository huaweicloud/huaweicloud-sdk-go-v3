package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordSetsResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]ListRecordSetsWithTags `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsResponse", string(data)}, " ")
}
