package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterLongAkskConfigResponse Response Object
type UpdateClusterLongAkskConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClusterLongAkskConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterLongAkskConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterLongAkskConfigResponse", string(data)}, " ")
}
