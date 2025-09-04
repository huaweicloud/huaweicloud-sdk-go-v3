package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFavouriteCustomTemplateResponse Response Object
type AddFavouriteCustomTemplateResponse struct {

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddFavouriteCustomTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFavouriteCustomTemplateResponse struct{}"
	}

	return strings.Join([]string{"AddFavouriteCustomTemplateResponse", string(data)}, " ")
}
