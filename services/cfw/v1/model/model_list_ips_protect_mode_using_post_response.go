package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsProtectModeUsingPostResponse Response Object
type ListIpsProtectModeUsingPostResponse struct {
	Data           *IpsProtectModeObject `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListIpsProtectModeUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsProtectModeUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ListIpsProtectModeUsingPostResponse", string(data)}, " ")
}
