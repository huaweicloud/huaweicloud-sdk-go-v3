package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessAgentLatestVersionResponse Response Object
type ListAccessAgentLatestVersionResponse struct {

	// 租户的HDP最新版本信息列表。
	Items          *[]LatestVersionInfo `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAccessAgentLatestVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessAgentLatestVersionResponse struct{}"
	}

	return strings.Join([]string{"ListAccessAgentLatestVersionResponse", string(data)}, " ")
}
