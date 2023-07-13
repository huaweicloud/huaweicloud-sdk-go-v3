package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePubInfoResponse Response Object
type CreatePubInfoResponse struct {
	Data           *CreatePubInfoResponseModel `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CreatePubInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePubInfoResponse struct{}"
	}

	return strings.Join([]string{"CreatePubInfoResponse", string(data)}, " ")
}
