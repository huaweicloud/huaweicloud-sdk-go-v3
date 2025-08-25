package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCcspClusterResponse Response Object
type DeleteCcspClusterResponse struct {

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteCcspClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCcspClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteCcspClusterResponse", string(data)}, " ")
}
