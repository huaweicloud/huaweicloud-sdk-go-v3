package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateApiResponse Response Object
type MigrateApiResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MigrateApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateApiResponse struct{}"
	}

	return strings.Join([]string{"MigrateApiResponse", string(data)}, " ")
}
