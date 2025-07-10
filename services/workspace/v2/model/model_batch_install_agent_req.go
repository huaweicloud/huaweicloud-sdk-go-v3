package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchInstallAgentReq 批量为桌面安装agent。
type BatchInstallAgentReq struct {

	// 操作的桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`
}

func (o BatchInstallAgentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInstallAgentReq struct{}"
	}

	return strings.Join([]string{"BatchInstallAgentReq", string(data)}, " ")
}
