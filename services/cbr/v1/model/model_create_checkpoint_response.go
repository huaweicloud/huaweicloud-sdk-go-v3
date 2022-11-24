package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCheckpointResponse struct {
	Checkpoint     *CheckpointCreate `json:"checkpoint,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CreateCheckpointResponse", string(data)}, " ")
}
