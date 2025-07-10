package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopPoolVolumesRequest Request Object
type AddDesktopPoolVolumesRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *AddDesktopPoolVolumesReq `json:"body,omitempty"`
}

func (o AddDesktopPoolVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopPoolVolumesRequest struct{}"
	}

	return strings.Join([]string{"AddDesktopPoolVolumesRequest", string(data)}, " ")
}
