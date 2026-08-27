package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlgorithmResponse Response Object
type UpdateAlgorithmResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateAlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlgorithmResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlgorithmResponse", string(data)}, " ")
}
