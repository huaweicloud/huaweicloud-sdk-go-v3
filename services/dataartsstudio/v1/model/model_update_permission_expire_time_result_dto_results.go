package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePermissionExpireTimeResultDtoResults struct {

	// 权限id
	Id *string `json:"id,omitempty"`

	// 更新有效期结果信息
	Message *string `json:"message,omitempty"`

	// 更新结果状态。1成功 0失败
	Status *int64 `json:"status,omitempty"`
}

func (o UpdatePermissionExpireTimeResultDtoResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionExpireTimeResultDtoResults struct{}"
	}

	return strings.Join([]string{"UpdatePermissionExpireTimeResultDtoResults", string(data)}, " ")
}
