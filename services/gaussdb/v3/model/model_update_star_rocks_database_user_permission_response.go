package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStarRocksDatabaseUserPermissionResponse Response Object
type UpdateStarRocksDatabaseUserPermissionResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStarRocksDatabaseUserPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarRocksDatabaseUserPermissionResponse struct{}"
	}

	return strings.Join([]string{"UpdateStarRocksDatabaseUserPermissionResponse", string(data)}, " ")
}
