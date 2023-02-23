package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterV2Response struct {
	Cluster        *Cluster `json:"cluster,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CreateClusterV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterV2Response struct{}"
	}

	return strings.Join([]string{"CreateClusterV2Response", string(data)}, " ")
}
