package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketsResponse Response Object
type ListBucketsResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListBucketsResponse", string(data)}, " ")
}
