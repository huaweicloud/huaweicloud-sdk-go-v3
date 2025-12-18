package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateToNodePoolResponse Response Object
type MigrateToNodePoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MigrateToNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateToNodePoolResponse struct{}"
	}

	return strings.Join([]string{"MigrateToNodePoolResponse", string(data)}, " ")
}
