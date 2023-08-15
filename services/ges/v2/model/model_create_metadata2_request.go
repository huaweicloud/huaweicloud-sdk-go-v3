package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetadata2Request Request Object
type CreateMetadata2Request struct {
	Body *CreateMetadataReq `json:"body,omitempty"`
}

func (o CreateMetadata2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadata2Request struct{}"
	}

	return strings.Join([]string{"CreateMetadata2Request", string(data)}, " ")
}
