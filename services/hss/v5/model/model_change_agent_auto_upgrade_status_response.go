package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAgentAutoUpgradeStatusResponse Response Object
type ChangeAgentAutoUpgradeStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeAgentAutoUpgradeStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAgentAutoUpgradeStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeAgentAutoUpgradeStatusResponse", string(data)}, " ")
}
