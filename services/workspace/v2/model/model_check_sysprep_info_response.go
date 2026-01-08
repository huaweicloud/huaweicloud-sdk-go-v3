package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSysprepInfoResponse Response Object
type CheckSysprepInfoResponse struct {

	// 桌面Sysprep信息。
	SysprepInfos   *[]DesktopSysprepInfo `json:"sysprep_infos,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CheckSysprepInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSysprepInfoResponse struct{}"
	}

	return strings.Join([]string{"CheckSysprepInfoResponse", string(data)}, " ")
}
