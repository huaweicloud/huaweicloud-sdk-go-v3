package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectionGroupResponse Response Object
type ShowProtectionGroupResponse struct {
	ServerGroup    *ShowProtectionGroupParams `json:"server_group,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowProtectionGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectionGroupResponse", string(data)}, " ")
}
