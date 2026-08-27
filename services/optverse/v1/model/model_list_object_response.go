package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObjectResponse Response Object
type ListObjectResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListObjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObjectResponse struct{}"
	}

	return strings.Join([]string{"ListObjectResponse", string(data)}, " ")
}
