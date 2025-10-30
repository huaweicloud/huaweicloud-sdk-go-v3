package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeAgentsResponse Response Object
type UpgradeAgentsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpgradeAgentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeAgentsResponse struct{}"
	}

	return strings.Join([]string{"UpgradeAgentsResponse", string(data)}, " ")
}
