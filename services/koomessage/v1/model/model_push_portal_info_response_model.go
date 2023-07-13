package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushPortalInfoResponseModel 主页催审返回体。
type PushPortalInfoResponseModel struct {
	Data *PushPortalInfoResponseModelData `json:"data,omitempty"`
}

func (o PushPortalInfoResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushPortalInfoResponseModel struct{}"
	}

	return strings.Join([]string{"PushPortalInfoResponseModel", string(data)}, " ")
}
