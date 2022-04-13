package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowReplicationResponse struct {
	Replication    *ShowReplicationParams `json:"replication,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationResponse", string(data)}, " ")
}
