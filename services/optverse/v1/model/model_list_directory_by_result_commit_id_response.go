package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectoryByResultCommitIdResponse Response Object
type ListDirectoryByResultCommitIdResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListDirectoryByResultCommitIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoryByResultCommitIdResponse struct{}"
	}

	return strings.Join([]string{"ListDirectoryByResultCommitIdResponse", string(data)}, " ")
}
