package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEvolveTaskResponse Response Object
type UpdateEvolveTaskResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateEvolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEvolveTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateEvolveTaskResponse", string(data)}, " ")
}
