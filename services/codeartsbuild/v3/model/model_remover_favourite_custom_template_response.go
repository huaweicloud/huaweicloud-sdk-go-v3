package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoverFavouriteCustomTemplateResponse Response Object
type RemoverFavouriteCustomTemplateResponse struct {

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoverFavouriteCustomTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoverFavouriteCustomTemplateResponse struct{}"
	}

	return strings.Join([]string{"RemoverFavouriteCustomTemplateResponse", string(data)}, " ")
}
