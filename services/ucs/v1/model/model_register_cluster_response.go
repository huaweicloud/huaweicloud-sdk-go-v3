package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterResponse Response Object
type RegisterClusterResponse struct {

	// 集群ID
	Uid            *string `json:"uid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterResponse struct{}"
	}

	return strings.Join([]string{"RegisterClusterResponse", string(data)}, " ")
}
