package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoverFavouriteCustomTemplateRequest Request Object
type RemoverFavouriteCustomTemplateRequest struct {

	// uuid
	Uuid string `json:"uuid"`
}

func (o RemoverFavouriteCustomTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoverFavouriteCustomTemplateRequest struct{}"
	}

	return strings.Join([]string{"RemoverFavouriteCustomTemplateRequest", string(data)}, " ")
}
