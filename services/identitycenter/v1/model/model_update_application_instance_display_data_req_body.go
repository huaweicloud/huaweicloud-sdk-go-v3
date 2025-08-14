package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceDisplayDataReqBody UpdateApplicationInstanceDisplayData的请求体
type UpdateApplicationInstanceDisplayDataReqBody struct {

	// 应用程序描述
	Description string `json:"description"`

	// 应用程序显示名
	DisplayName string `json:"display_name"`
}

func (o UpdateApplicationInstanceDisplayDataReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceDisplayDataReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceDisplayDataReqBody", string(data)}, " ")
}
