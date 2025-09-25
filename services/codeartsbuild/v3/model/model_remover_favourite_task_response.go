package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoverFavouriteTaskResponse Response Object
type RemoverFavouriteTaskResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *FavouriteResponseResult `json:"result,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o RemoverFavouriteTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoverFavouriteTaskResponse struct{}"
	}

	return strings.Join([]string{"RemoverFavouriteTaskResponse", string(data)}, " ")
}
