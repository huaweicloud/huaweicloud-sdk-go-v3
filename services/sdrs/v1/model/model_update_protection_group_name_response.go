package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateProtectionGroupNameResponse struct {
	ServerGroup    *ShowProtectionGroupParams `json:"server_group,omitempty" xml:"server_group"`
	HttpStatusCode int                        `json:"-"`
}

func (o UpdateProtectionGroupNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionGroupNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectionGroupNameResponse", string(data)}, " ")
}
