package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStarRocksDatabaseUserPasswordResponse Response Object
type UpdateStarRocksDatabaseUserPasswordResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStarRocksDatabaseUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarRocksDatabaseUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdateStarRocksDatabaseUserPasswordResponse", string(data)}, " ")
}
