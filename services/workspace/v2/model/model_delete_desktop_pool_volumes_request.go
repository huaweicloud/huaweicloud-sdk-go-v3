package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopPoolVolumesRequest Request Object
type DeleteDesktopPoolVolumesRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *DeleteDesktopPoolVolumesReq `json:"body,omitempty"`
}

func (o DeleteDesktopPoolVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolVolumesRequest struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolVolumesRequest", string(data)}, " ")
}
