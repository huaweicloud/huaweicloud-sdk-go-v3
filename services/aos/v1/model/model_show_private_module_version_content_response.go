package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateModuleVersionContentResponse Response Object
type ShowPrivateModuleVersionContentResponse struct {
	Location       *string `json:"Location,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPrivateModuleVersionContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateModuleVersionContentResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateModuleVersionContentResponse", string(data)}, " ")
}
