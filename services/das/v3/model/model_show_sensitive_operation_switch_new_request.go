package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSensitiveOperationSwitchNewRequest Request Object
type ShowSensitiveOperationSwitchNewRequest struct {
}

func (o ShowSensitiveOperationSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSensitiveOperationSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowSensitiveOperationSwitchNewRequest", string(data)}, " ")
}
