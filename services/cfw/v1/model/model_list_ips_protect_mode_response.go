package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsProtectModeResponse Response Object
type ListIpsProtectModeResponse struct {
	Data           *IpsProtectModeObject `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListIpsProtectModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsProtectModeResponse struct{}"
	}

	return strings.Join([]string{"ListIpsProtectModeResponse", string(data)}, " ")
}
