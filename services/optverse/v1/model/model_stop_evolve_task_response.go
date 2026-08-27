package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopEvolveTaskResponse Response Object
type StopEvolveTaskResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o StopEvolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopEvolveTaskResponse struct{}"
	}

	return strings.Join([]string{"StopEvolveTaskResponse", string(data)}, " ")
}
