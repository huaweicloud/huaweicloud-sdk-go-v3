package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordSetsByZoneResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]ListRecordSets `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRecordSetsByZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsByZoneResponse struct{}"
	}

	return strings.Join([]string{"ListRecordSetsByZoneResponse", string(data)}, " ")
}
