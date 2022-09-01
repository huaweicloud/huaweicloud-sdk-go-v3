package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateClusterTagRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *CreateTagReq `json:"body,omitempty" xml:"body"`
}

func (o CreateClusterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterTagRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterTagRequest", string(data)}, " ")
}
