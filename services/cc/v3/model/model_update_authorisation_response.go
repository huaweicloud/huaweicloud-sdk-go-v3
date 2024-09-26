package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthorisationResponse Response Object
type UpdateAuthorisationResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	Authorisation  *Authorisation `json:"authorisation"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateAuthorisationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthorisationResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuthorisationResponse", string(data)}, " ")
}
