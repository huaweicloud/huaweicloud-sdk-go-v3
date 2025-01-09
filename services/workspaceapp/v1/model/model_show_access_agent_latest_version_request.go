package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessAgentLatestVersionRequest Request Object
type ShowAccessAgentLatestVersionRequest struct {
}

func (o ShowAccessAgentLatestVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessAgentLatestVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessAgentLatestVersionRequest", string(data)}, " ")
}
