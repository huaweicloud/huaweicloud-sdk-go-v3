package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageTypeRiskInfo struct {

	// 本地镜像
	Local *int32 `json:"local,omitempty"`

	// 仓库镜像
	Registriy *int32 `json:"registriy,omitempty"`

	// cicd镜像
	Cicd *int32 `json:"cicd,omitempty"`
}

func (o ImageTypeRiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTypeRiskInfo struct{}"
	}

	return strings.Join([]string{"ImageTypeRiskInfo", string(data)}, " ")
}
