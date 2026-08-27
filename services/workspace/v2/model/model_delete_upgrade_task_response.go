package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUpgradeTaskResponse Response Object
type DeleteUpgradeTaskResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteUpgradeTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUpgradeTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteUpgradeTaskResponse", string(data)}, " ")
}
