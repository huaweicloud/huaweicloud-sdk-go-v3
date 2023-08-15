package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceTopicResponse Response Object
type UpdateInstanceTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicResponse", string(data)}, " ")
}
