package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPtrRecordsResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Metadata *Metedata `json:"metadata,omitempty" xml:"metadata"`

	Floatingips    *[]ListPtrRecordsFloatingResp `json:"floatingips,omitempty" xml:"floatingips"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPtrRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPtrRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPtrRecordsResponse", string(data)}, " ")
}
