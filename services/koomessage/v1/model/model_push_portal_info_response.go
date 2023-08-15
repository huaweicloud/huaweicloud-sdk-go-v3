package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushPortalInfoResponse Response Object
type PushPortalInfoResponse struct {
	Data           *PushPortalInfoResponseModel `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o PushPortalInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushPortalInfoResponse struct{}"
	}

	return strings.Join([]string{"PushPortalInfoResponse", string(data)}, " ")
}
