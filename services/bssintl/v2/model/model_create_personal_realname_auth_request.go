package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePersonalRealnameAuthRequest Request Object
type CreatePersonalRealnameAuthRequest struct {
	Body *ApplyIndividualRealnameAuthsReq `json:"body,omitempty"`
}

func (o CreatePersonalRealnameAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersonalRealnameAuthRequest struct{}"
	}

	return strings.Join([]string{"CreatePersonalRealnameAuthRequest", string(data)}, " ")
}
