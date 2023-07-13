package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicResponse Response Object
type UpdateTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicResponse", string(data)}, " ")
}
