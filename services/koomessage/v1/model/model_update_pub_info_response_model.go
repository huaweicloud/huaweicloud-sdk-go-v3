package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePubInfoResponseModel 更改服务号详情返回体。
type UpdatePubInfoResponseModel struct {
	Data *UpdatePubInfoResponseModelData `json:"data,omitempty"`
}

func (o UpdatePubInfoResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePubInfoResponseModel struct{}"
	}

	return strings.Join([]string{"UpdatePubInfoResponseModel", string(data)}, " ")
}
