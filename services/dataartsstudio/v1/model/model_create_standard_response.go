package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardResponse Response Object
type CreateStandardResponse struct {
	Data           *CreateStandardResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CreateStandardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardResponse struct{}"
	}

	return strings.Join([]string{"CreateStandardResponse", string(data)}, " ")
}
