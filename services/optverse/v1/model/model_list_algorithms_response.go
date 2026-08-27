package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlgorithmsResponse Response Object
type ListAlgorithmsResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAlgorithmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlgorithmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlgorithmsResponse", string(data)}, " ")
}
