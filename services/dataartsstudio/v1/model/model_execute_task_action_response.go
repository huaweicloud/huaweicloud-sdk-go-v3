package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTaskActionResponse Response Object
type ExecuteTaskActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteTaskActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTaskActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTaskActionResponse", string(data)}, " ")
}
