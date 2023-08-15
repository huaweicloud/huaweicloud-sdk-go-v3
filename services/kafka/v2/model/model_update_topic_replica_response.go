package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicReplicaResponse Response Object
type UpdateTopicReplicaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicReplicaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicReplicaResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicReplicaResponse", string(data)}, " ")
}
