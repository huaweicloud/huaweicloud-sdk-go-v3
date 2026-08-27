package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEvolveTaskResponse Response Object
type CreateEvolveTaskResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateEvolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEvolveTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateEvolveTaskResponse", string(data)}, " ")
}
