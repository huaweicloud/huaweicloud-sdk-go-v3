package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagRequest Request Object
type DeleteTagRequest struct {

	// 桌面id。
	DesktopId string `json:"desktop_id"`

	// 标签key。
	Key string `json:"key"`
}

func (o DeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}
