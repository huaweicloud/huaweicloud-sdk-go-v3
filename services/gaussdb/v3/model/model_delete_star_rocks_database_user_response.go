package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRocksDatabaseUserResponse Response Object
type DeleteStarRocksDatabaseUserResponse struct {

	// 请求结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteStarRocksDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRocksDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteStarRocksDatabaseUserResponse", string(data)}, " ")
}
