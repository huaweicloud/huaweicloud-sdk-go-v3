package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopRsp 创建按需桌面时的响应体，可根据job_id查询资源是否创建成功。
type CreateDesktopRsp struct {

	// 创建云桌面总任务id
	JobId *string `json:"job_id,omitempty"`
}

func (o CreateDesktopRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopRsp struct{}"
	}

	return strings.Join([]string{"CreateDesktopRsp", string(data)}, " ")
}
