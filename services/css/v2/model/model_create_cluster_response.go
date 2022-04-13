package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterResponse struct {
	// 集群ID。

	Id *string `json:"id,omitempty"`
	// 集群名称。

	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
