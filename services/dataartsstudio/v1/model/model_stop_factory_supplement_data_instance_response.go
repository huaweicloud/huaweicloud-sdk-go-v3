package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFactorySupplementDataInstanceResponse Response Object
type StopFactorySupplementDataInstanceResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopFactorySupplementDataInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFactorySupplementDataInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopFactorySupplementDataInstanceResponse", string(data)}, " ")
}
