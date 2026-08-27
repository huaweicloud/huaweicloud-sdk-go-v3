package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEvolveTaskResponse Response Object
type DeleteEvolveTaskResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DeleteEvolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvolveTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteEvolveTaskResponse", string(data)}, " ")
}
