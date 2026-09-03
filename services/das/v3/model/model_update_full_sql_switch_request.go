package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFullSqlSwitchRequest Request Object
type UpdateFullSqlSwitchRequest struct {
	Body *UpdateFullSqlSwitchRequestBody `json:"body,omitempty"`
}

func (o UpdateFullSqlSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFullSqlSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateFullSqlSwitchRequest", string(data)}, " ")
}
