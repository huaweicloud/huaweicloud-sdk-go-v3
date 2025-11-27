package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableGitOpsRequest Request Object
type DisableGitOpsRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`
}

func (o DisableGitOpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableGitOpsRequest struct{}"
	}

	return strings.Join([]string{"DisableGitOpsRequest", string(data)}, " ")
}
