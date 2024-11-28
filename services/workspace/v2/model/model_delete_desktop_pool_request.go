package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopPoolRequest Request Object
type DeleteDesktopPoolRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`
}

func (o DeleteDesktopPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolRequest struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolRequest", string(data)}, " ")
}
