package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopPoolDetailRequest Request Object
type ShowDesktopPoolDetailRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`
}

func (o ShowDesktopPoolDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopPoolDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopPoolDetailRequest", string(data)}, " ")
}
