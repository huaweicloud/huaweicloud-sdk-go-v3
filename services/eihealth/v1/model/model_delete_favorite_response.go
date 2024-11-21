package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFavoriteResponse Response Object
type DeleteFavoriteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFavoriteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFavoriteResponse struct{}"
	}

	return strings.Join([]string{"DeleteFavoriteResponse", string(data)}, " ")
}
