package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRecordSetByZoneResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Recordsets *[]ShowRecordSetByZoneResp `json:"recordsets,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowRecordSetByZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordSetByZoneResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordSetByZoneResponse", string(data)}, " ")
}
