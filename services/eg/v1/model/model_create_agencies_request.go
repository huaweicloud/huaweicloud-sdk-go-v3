package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgenciesRequest Request Object
type CreateAgenciesRequest struct {
	Body *AgencyCreateReq `json:"body,omitempty"`
}

func (o CreateAgenciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesRequest struct{}"
	}

	return strings.Join([]string{"CreateAgenciesRequest", string(data)}, " ")
}
