package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterConfRequest Request Object
type CreateClusterConfRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`

	Body *RequiredInput `json:"body,omitempty"`
}

func (o CreateClusterConfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterConfRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterConfRequest", string(data)}, " ")
}
