package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIncidentUserInfo struct {

	// Id value
	UserId *string `json:"user_id,omitempty"`

	// The name, display only
	UserName *string `json:"user_name,omitempty"`
}

func (o CreateIncidentUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncidentUserInfo struct{}"
	}

	return strings.Join([]string{"CreateIncidentUserInfo", string(data)}, " ")
}
