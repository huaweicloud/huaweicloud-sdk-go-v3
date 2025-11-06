package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePermissionBodyDto struct {

	// **参数解释：** 更新资源权限详情 **取值范围：** 不涉及。
	Data *[]UpdatePermissionDto `json:"data,omitempty"`
}

func (o UpdatePermissionBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionBodyDto struct{}"
	}

	return strings.Join([]string{"UpdatePermissionBodyDto", string(data)}, " ")
}
