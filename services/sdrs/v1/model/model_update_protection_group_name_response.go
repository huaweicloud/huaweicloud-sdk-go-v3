package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProtectionGroupNameResponse Response Object
type UpdateProtectionGroupNameResponse struct {
	ServerGroup    *ShowProtectionGroupParams `json:"server_group,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o UpdateProtectionGroupNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionGroupNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectionGroupNameResponse", string(data)}, " ")
}
