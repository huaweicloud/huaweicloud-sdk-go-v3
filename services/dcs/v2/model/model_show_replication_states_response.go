package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplicationStatesResponse Response Object
type ShowReplicationStatesResponse struct {
	Body           *[]InstanceReplicationListInfo `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowReplicationStatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationStatesResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationStatesResponse", string(data)}, " ")
}
