package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlgorithmResponse Response Object
type CreateAlgorithmResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateAlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlgorithmResponse struct{}"
	}

	return strings.Join([]string{"CreateAlgorithmResponse", string(data)}, " ")
}
