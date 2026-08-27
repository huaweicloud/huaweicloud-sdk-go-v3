package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteEvolveTaskResponse Response Object
type BatchDeleteEvolveTaskResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchDeleteEvolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteEvolveTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteEvolveTaskResponse", string(data)}, " ")
}
