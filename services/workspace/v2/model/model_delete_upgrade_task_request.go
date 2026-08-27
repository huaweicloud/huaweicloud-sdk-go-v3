package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUpgradeTaskRequest Request Object
type DeleteUpgradeTaskRequest struct {
	Body *DeleteUpgradeTaskRequestBody `json:"body,omitempty"`
}

func (o DeleteUpgradeTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUpgradeTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteUpgradeTaskRequest", string(data)}, " ")
}
