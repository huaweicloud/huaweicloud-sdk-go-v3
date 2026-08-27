package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAlgorithmFileResponse Response Object
type ImportAlgorithmFileResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ImportAlgorithmFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAlgorithmFileResponse struct{}"
	}

	return strings.Join([]string{"ImportAlgorithmFileResponse", string(data)}, " ")
}
