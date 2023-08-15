package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentByNameResponse Response Object
type ShowComponentByNameResponse struct {

	// 组件名称
	Name *string `json:"name,omitempty"`

	// 组件id
	Id *string `json:"id,omitempty"`

	// aomId
	AomId *string `json:"aom_id,omitempty"`

	// 应用id
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowComponentByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentByNameResponse struct{}"
	}

	return strings.Join([]string{"ShowComponentByNameResponse", string(data)}, " ")
}
