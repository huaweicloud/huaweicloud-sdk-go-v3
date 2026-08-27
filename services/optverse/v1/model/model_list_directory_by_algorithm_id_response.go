package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectoryByAlgorithmIdResponse Response Object
type ListDirectoryByAlgorithmIdResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListDirectoryByAlgorithmIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoryByAlgorithmIdResponse struct{}"
	}

	return strings.Join([]string{"ListDirectoryByAlgorithmIdResponse", string(data)}, " ")
}
