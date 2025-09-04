package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteServerDumpResponse Response Object
type ExecuteServerDumpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteServerDumpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteServerDumpResponse struct{}"
	}

	return strings.Join([]string{"ExecuteServerDumpResponse", string(data)}, " ")
}
