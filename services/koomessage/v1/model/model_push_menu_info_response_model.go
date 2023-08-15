package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushMenuInfoResponseModel 菜单催审返回体。
type PushMenuInfoResponseModel struct {
	Data *PushMenuInfoResponseModelData `json:"data,omitempty"`
}

func (o PushMenuInfoResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushMenuInfoResponseModel struct{}"
	}

	return strings.Join([]string{"PushMenuInfoResponseModel", string(data)}, " ")
}
