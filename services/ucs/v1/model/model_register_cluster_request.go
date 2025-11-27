package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterRequest Request Object
type RegisterClusterRequest struct {
	Body *RegisterCluster `json:"body,omitempty" type:"multipart"`
}

func (o RegisterClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterRequest struct{}"
	}

	return strings.Join([]string{"RegisterClusterRequest", string(data)}, " ")
}
