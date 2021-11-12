package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CommitCheckpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CommitCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CommitCheckpointResponse", string(data)}, " ")
}
