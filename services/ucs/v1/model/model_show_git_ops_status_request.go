package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGitOpsStatusRequest Request Object
type ShowGitOpsStatusRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`
}

func (o ShowGitOpsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGitOpsStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowGitOpsStatusRequest", string(data)}, " ")
}
