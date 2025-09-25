package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFavouriteOfficialTemplateRequest Request Object
type AddFavouriteOfficialTemplateRequest struct {

	// uuid
	Uuid string `json:"uuid"`
}

func (o AddFavouriteOfficialTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFavouriteOfficialTemplateRequest struct{}"
	}

	return strings.Join([]string{"AddFavouriteOfficialTemplateRequest", string(data)}, " ")
}
