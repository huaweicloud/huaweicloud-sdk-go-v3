package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePubInfoResponse Response Object
type UpdatePubInfoResponse struct {
	Data           *UpdatePubInfoResponseModel `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdatePubInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePubInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdatePubInfoResponse", string(data)}, " ")
}
