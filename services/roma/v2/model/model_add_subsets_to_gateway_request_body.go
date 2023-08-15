package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSubsetsToGatewayRequestBody struct {

	// 待添加子设备ID列表，且设备需是普通设备，自动向下取整
	Resources []int32 `json:"resources"`
}

func (o AddSubsetsToGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubsetsToGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"AddSubsetsToGatewayRequestBody", string(data)}, " ")
}
