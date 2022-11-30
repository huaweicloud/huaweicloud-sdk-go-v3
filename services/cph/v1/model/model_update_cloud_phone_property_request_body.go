package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新云手机属性请求体。
type UpdateCloudPhonePropertyRequestBody struct {

	// 手机列表
	Phones []Property `json:"phones"`
}

func (o UpdateCloudPhonePropertyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudPhonePropertyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCloudPhonePropertyRequestBody", string(data)}, " ")
}
