package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClickHouseDatabaseUserPermissionResponse Response Object
type UpdateClickHouseDatabaseUserPermissionResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClickHouseDatabaseUserPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClickHouseDatabaseUserPermissionResponse struct{}"
	}

	return strings.Join([]string{"UpdateClickHouseDatabaseUserPermissionResponse", string(data)}, " ")
}
