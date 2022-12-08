package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWorkspaceLockResponse struct {

	// 云办公服务是否被锁定，0代表未锁定，1代表锁定。
	IsLock *int32 `json:"is_lock,omitempty"`

	// 云办公服务锁定时间，格式：yyyy-MM-dd HH:mm:ss，时区：UTC。
	LockTime *string `json:"lock_time,omitempty"`

	// 云办公服务锁定原因。
	LockReason     *string `json:"lock_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWorkspaceLockResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceLockResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceLockResponse", string(data)}, " ")
}
