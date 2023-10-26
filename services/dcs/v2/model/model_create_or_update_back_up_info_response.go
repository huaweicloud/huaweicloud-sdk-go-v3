package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateBackUpInfoResponse Response Object
type CreateOrUpdateBackUpInfoResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateOrUpdateBackUpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateBackUpInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateBackUpInfoResponse", string(data)}, " ")
}
