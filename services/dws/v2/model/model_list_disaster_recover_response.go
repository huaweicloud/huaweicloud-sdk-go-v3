package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDisasterRecoverResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDisasterRecoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoverResponse struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoverResponse", string(data)}, " ")
}
