package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlgorithmFileResponse Response Object
type ShowAlgorithmFileResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowAlgorithmFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlgorithmFileResponse struct{}"
	}

	return strings.Join([]string{"ShowAlgorithmFileResponse", string(data)}, " ")
}
