package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferIntranetConnectionRequest Request Object
type CreateInferIntranetConnectionRequest struct {
	Body *IntranetConnectionRequest `json:"body,omitempty"`
}

func (o CreateInferIntranetConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferIntranetConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateInferIntranetConnectionRequest", string(data)}, " ")
}
