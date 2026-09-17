package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSensitiveOperationSwitchNewResponse Response Object
type ShowSensitiveOperationSwitchNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowSensitiveOperationSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSensitiveOperationSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"ShowSensitiveOperationSwitchNewResponse", string(data)}, " ")
}
