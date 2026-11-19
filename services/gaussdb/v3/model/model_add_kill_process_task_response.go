package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddKillProcessTaskResponse Response Object
type AddKillProcessTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddKillProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddKillProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"AddKillProcessTaskResponse", string(data)}, " ")
}
