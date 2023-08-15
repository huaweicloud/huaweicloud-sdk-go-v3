package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCheckpointResponse Response Object
type ShowCheckpointResponse struct {
	Checkpoint     *CheckpointCreate `json:"checkpoint,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCheckpointResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckpointResponse", string(data)}, " ")
}
