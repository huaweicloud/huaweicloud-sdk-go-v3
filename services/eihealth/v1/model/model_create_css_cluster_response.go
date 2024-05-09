package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCssClusterResponse Response Object
type CreateCssClusterResponse struct {

	// css集群id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCssClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCssClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateCssClusterResponse", string(data)}, " ")
}
