package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRawIconReq 更新应用图标
type UpdateRawIconReq struct {

	// 待更新的应用图标
	IconContent string `json:"icon_content"`
}

func (o UpdateRawIconReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRawIconReq struct{}"
	}

	return strings.Join([]string{"UpdateRawIconReq", string(data)}, " ")
}
