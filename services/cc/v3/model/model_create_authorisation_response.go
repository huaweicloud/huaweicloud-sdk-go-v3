package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorisationResponse Response Object
type CreateAuthorisationResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	Authorisation  *Authorisation `json:"authorisation"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateAuthorisationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorisationResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthorisationResponse", string(data)}, " ")
}
