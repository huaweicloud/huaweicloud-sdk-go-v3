package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordSetsByZoneResponse struct {
	Links *PageLink `json:"links,omitempty" xml:"links"`

	Recordsets *[]ListRecordSets `json:"recordsets,omitempty" xml:"recordsets"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsByZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsByZoneResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsByZoneResponse", string(data)}, " ")
}
