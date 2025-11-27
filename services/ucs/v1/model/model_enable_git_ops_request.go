package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableGitOpsRequest Request Object
type EnableGitOpsRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`

	// 重试参数
	Retry *string `json:"retry,omitempty"`
}

func (o EnableGitOpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableGitOpsRequest struct{}"
	}

	return strings.Join([]string{"EnableGitOpsRequest", string(data)}, " ")
}
