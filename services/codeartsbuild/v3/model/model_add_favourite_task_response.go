package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFavouriteTaskResponse Response Object
type AddFavouriteTaskResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *FavouriteResponseResult `json:"result,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o AddFavouriteTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFavouriteTaskResponse struct{}"
	}

	return strings.Join([]string{"AddFavouriteTaskResponse", string(data)}, " ")
}
