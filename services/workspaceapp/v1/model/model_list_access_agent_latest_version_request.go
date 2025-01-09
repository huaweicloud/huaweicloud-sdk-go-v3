package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessAgentLatestVersionRequest Request Object
type ListAccessAgentLatestVersionRequest struct {
}

func (o ListAccessAgentLatestVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessAgentLatestVersionRequest struct{}"
	}

	return strings.Join([]string{"ListAccessAgentLatestVersionRequest", string(data)}, " ")
}
