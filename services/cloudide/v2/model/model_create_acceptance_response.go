package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAcceptanceResponse Response Object
type CreateAcceptanceResponse struct {

	// acceptance_id
	AcceptanceId   *int32 `json:"acceptance_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateAcceptanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceptanceResponse struct{}"
	}

	return strings.Join([]string{"CreateAcceptanceResponse", string(data)}, " ")
}
