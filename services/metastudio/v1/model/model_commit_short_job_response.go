package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitShortJobResponse Response Object
type CommitShortJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CommitShortJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitShortJobResponse struct{}"
	}

	return strings.Join([]string{"CommitShortJobResponse", string(data)}, " ")
}
