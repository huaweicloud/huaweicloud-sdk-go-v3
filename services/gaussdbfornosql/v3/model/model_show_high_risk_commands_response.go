package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHighRiskCommandsResponse Response Object
type ShowHighRiskCommandsResponse struct {

	// 高危命令与对应重命名命令。
	Commands       *[]CommandInfo `json:"commands,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHighRiskCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHighRiskCommandsResponse struct{}"
	}

	return strings.Join([]string{"ShowHighRiskCommandsResponse", string(data)}, " ")
}
