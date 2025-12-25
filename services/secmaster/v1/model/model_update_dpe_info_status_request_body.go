package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDpeInfoStatusRequestBody 更新分类映射状态请求体
type UpdateDpeInfoStatusRequestBody struct {

	// 状态(enabled，disabled)
	Status string `json:"status"`
}

func (o UpdateDpeInfoStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDpeInfoStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDpeInfoStatusRequestBody", string(data)}, " ")
}
