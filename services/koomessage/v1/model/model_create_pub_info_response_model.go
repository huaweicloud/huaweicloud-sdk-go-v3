package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePubInfoResponseModel 创建服务号的返回体。
type CreatePubInfoResponseModel struct {
	Data *CreatePubInfoResponseModelData `json:"data,omitempty"`
}

func (o CreatePubInfoResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePubInfoResponseModel struct{}"
	}

	return strings.Join([]string{"CreatePubInfoResponseModel", string(data)}, " ")
}
