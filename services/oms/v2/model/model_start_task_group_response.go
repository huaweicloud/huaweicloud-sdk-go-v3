package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTaskGroupResponse Response Object
type StartTaskGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartTaskGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTaskGroupResponse struct{}"
	}

	return strings.Join([]string{"StartTaskGroupResponse", string(data)}, " ")
}
