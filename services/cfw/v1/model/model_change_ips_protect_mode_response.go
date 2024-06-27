package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIpsProtectModeResponse Response Object
type ChangeIpsProtectModeResponse struct {
	Data           *CommonResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ChangeIpsProtectModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsProtectModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeIpsProtectModeResponse", string(data)}, " ")
}
