package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFavouriteOfficialTemplateResponse Response Object
type AddFavouriteOfficialTemplateResponse struct {

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddFavouriteOfficialTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFavouriteOfficialTemplateResponse struct{}"
	}

	return strings.Join([]string{"AddFavouriteOfficialTemplateResponse", string(data)}, " ")
}
