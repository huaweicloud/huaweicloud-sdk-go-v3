package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAkSksResponse struct {
	AccessAkSkModels *[]AccessAkskVo `json:"access_ak_sk_models,omitempty"`
	HttpStatusCode   int             `json:"-"`
}

func (o ShowAkSksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAkSksResponse struct{}"
	}

	return strings.Join([]string{"ShowAkSksResponse", string(data)}, " ")
}
