package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFavouriteCustomTemplateRequest Request Object
type AddFavouriteCustomTemplateRequest struct {

	// uuid
	Uuid string `json:"uuid"`
}

func (o AddFavouriteCustomTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFavouriteCustomTemplateRequest struct{}"
	}

	return strings.Join([]string{"AddFavouriteCustomTemplateRequest", string(data)}, " ")
}
