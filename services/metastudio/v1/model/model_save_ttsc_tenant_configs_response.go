package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTtscTenantConfigsResponse Response Object
type SaveTtscTenantConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SaveTtscTenantConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTtscTenantConfigsResponse struct{}"
	}

	return strings.Join([]string{"SaveTtscTenantConfigsResponse", string(data)}, " ")
}
