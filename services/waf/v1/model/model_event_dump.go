package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventDump struct {
	Source *string `json:"source,omitempty"`

	State *string `json:"state,omitempty"`

	Description *string `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	Filename *string `json:"filename,omitempty"`

	Obsname *string `json:"obsname,omitempty"`

	Tenantid *string `json:"tenantid,omitempty"`

	Start *int64 `json:"start,omitempty"`

	End *int64 `json:"end,omitempty"`

	Total *int64 `json:"total,omitempty"`

	Url *string `json:"url,omitempty"`

	Urltimestamp *int64 `json:"urltimestamp,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o EventDump) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDump struct{}"
	}

	return strings.Join([]string{"EventDump", string(data)}, " ")
}
