package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EdgeAppInstanceDto struct {

	// 边缘应用id，只允许数字、英文小写、中划线，切必须以字母或数字结尾
	EdgeAppId string `json:"edge_app_id" xml:"edge_app_id"`

	// 边缘应用版本，只允许数字、英文小写、中划线，切必须以字母或数字结尾
	AppVersion string `json:"app_version" xml:"app_version"`
}

func (o EdgeAppInstanceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeAppInstanceDto struct{}"
	}

	return strings.Join([]string{"EdgeAppInstanceDto", string(data)}, " ")
}
