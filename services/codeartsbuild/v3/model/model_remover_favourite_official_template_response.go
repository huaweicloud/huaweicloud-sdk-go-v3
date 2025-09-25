package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoverFavouriteOfficialTemplateResponse Response Object
type RemoverFavouriteOfficialTemplateResponse struct {

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoverFavouriteOfficialTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoverFavouriteOfficialTemplateResponse struct{}"
	}

	return strings.Join([]string{"RemoverFavouriteOfficialTemplateResponse", string(data)}, " ")
}
