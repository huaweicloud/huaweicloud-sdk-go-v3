package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootClusterResponse Response Object
type RebootClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RebootClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootClusterResponse struct{}"
	}

	return strings.Join([]string{"RebootClusterResponse", string(data)}, " ")
}
