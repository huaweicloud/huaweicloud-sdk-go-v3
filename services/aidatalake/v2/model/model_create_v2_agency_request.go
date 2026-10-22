package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2AgencyRequest Request Object
type CreateV2AgencyRequest struct {
	Body *CreateV2AgencyRequestBody `json:"body,omitempty"`
}

func (o CreateV2AgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2AgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateV2AgencyRequest", string(data)}, " ")
}
