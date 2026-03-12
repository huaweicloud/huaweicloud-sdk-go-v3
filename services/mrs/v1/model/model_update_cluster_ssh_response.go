package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterSshResponse Response Object
type UpdateClusterSshResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClusterSshResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterSshResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterSshResponse", string(data)}, " ")
}
