package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommandResponse Response Object
type ShowCommandResponse struct {

	// 命令名称，分为：START、STOP、SKIP、PREMIGRATE、TASK_NOT_STARTED、TASK_HAS_STOPPED、TASK_HAS_FINISHED
	CommandName *string `json:"command_name,omitempty"`

	CommandParam   *CommandParam `json:"command_param,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommandResponse struct{}"
	}

	return strings.Join([]string{"ShowCommandResponse", string(data)}, " ")
}
