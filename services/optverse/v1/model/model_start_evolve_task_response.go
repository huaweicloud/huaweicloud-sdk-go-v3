package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartEvolveTaskResponse Response Object
type StartEvolveTaskResponse struct {
	MetaInfo *MetaInfo `json:"meta_info,omitempty"`

	Payload        *PayloadObject `json:"payload,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o StartEvolveTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartEvolveTaskResponse struct{}"
	}

	return strings.Join([]string{"StartEvolveTaskResponse", string(data)}, " ")
}
