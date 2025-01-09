package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessAgentLatestVersionResponse Response Object
type ShowAccessAgentLatestVersionResponse struct {

	// 租户的HDA最新版本。
	LatestVersion  *string `json:"latest_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAccessAgentLatestVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessAgentLatestVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessAgentLatestVersionResponse", string(data)}, " ")
}
