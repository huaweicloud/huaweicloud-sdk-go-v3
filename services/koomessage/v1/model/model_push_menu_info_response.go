package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushMenuInfoResponse Response Object
type PushMenuInfoResponse struct {
	Data           *PushMenuInfoResponseModel `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o PushMenuInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushMenuInfoResponse struct{}"
	}

	return strings.Join([]string{"PushMenuInfoResponse", string(data)}, " ")
}
