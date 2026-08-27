package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveAlgorithmFileResponse Response Object
type SaveAlgorithmFileResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o SaveAlgorithmFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAlgorithmFileResponse struct{}"
	}

	return strings.Join([]string{"SaveAlgorithmFileResponse", string(data)}, " ")
}
