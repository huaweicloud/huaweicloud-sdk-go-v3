package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckClusterRequest struct {
	Body *ClusterCheckRequestBody `json:"body,omitempty"`
}

func (o CheckClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClusterRequest struct{}"
	}

	return strings.Join([]string{"CheckClusterRequest", string(data)}, " ")
}
