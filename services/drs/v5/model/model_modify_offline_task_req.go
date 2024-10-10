package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyOfflineTaskReq 修改离线迁移任务请求。
type ModifyOfflineTaskReq struct {

	// 备份迁移任务名称。
	Name string `json:"name"`

	// 备份迁移任务描述。
	Description *string `json:"description,omitempty"`
}

func (o ModifyOfflineTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyOfflineTaskReq struct{}"
	}

	return strings.Join([]string{"ModifyOfflineTaskReq", string(data)}, " ")
}
