package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFavoriteResponse Response Object
type CreateFavoriteResponse struct {

	// 收藏记录ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFavoriteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFavoriteResponse struct{}"
	}

	return strings.Join([]string{"CreateFavoriteResponse", string(data)}, " ")
}
