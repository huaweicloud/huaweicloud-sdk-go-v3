package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateAzResponse Response Object
type MigrateAzResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MigrateAzResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzResponse struct{}"
	}

	return strings.Join([]string{"MigrateAzResponse", string(data)}, " ")
}
