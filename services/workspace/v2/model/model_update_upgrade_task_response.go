package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUpgradeTaskResponse Response Object
type UpdateUpgradeTaskResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateUpgradeTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUpgradeTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateUpgradeTaskResponse", string(data)}, " ")
}
