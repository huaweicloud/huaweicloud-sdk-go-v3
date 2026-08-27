package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlgorithmFileResponse Response Object
type DeleteAlgorithmFileResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DeleteAlgorithmFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlgorithmFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlgorithmFileResponse", string(data)}, " ")
}
