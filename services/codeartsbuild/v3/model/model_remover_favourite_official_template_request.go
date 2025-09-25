package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoverFavouriteOfficialTemplateRequest Request Object
type RemoverFavouriteOfficialTemplateRequest struct {

	// uuid
	Uuid string `json:"uuid"`
}

func (o RemoverFavouriteOfficialTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoverFavouriteOfficialTemplateRequest struct{}"
	}

	return strings.Join([]string{"RemoverFavouriteOfficialTemplateRequest", string(data)}, " ")
}
