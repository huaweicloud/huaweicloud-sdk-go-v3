package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactorySupplementDataInstanceResponse Response Object
type CreateFactorySupplementDataInstanceResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFactorySupplementDataInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactorySupplementDataInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateFactorySupplementDataInstanceResponse", string(data)}, " ")
}
