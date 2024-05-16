package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStarRocksDatabaseUserResponse Response Object
type CreateStarRocksDatabaseUserResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStarRocksDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStarRocksDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"CreateStarRocksDatabaseUserResponse", string(data)}, " ")
}
