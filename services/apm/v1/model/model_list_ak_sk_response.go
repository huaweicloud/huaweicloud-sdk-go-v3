package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAkSkResponse struct {
	AccessAkSkModels *[]AccessAkskVo `json:"access_ak_sk_models,omitempty"`
	HttpStatusCode   int             `json:"-"`
}

func (o ListAkSkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAkSkResponse struct{}"
	}

	return strings.Join([]string{"ListAkSkResponse", string(data)}, " ")
}
