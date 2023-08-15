package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopData 变更规格的桌面对象。
type ResizeDesktopData struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`
}

func (o ResizeDesktopData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopData struct{}"
	}

	return strings.Join([]string{"ResizeDesktopData", string(data)}, " ")
}
