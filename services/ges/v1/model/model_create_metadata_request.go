package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMetadataRequest struct {
	Body *CreateMetadataReq `json:"body,omitempty"`
}

func (o CreateMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataRequest struct{}"
	}

	return strings.Join([]string{"CreateMetadataRequest", string(data)}, " ")
}
