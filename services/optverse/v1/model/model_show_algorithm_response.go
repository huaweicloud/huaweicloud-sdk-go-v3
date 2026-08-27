package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlgorithmResponse Response Object
type ShowAlgorithmResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowAlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlgorithmResponse struct{}"
	}

	return strings.Join([]string{"ShowAlgorithmResponse", string(data)}, " ")
}
