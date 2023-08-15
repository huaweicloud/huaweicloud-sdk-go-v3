package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreClusterResponse Response Object
type RestoreClusterResponse struct {
	Cluster        *Cluster `json:"cluster,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o RestoreClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreClusterResponse struct{}"
	}

	return strings.Join([]string{"RestoreClusterResponse", string(data)}, " ")
}
