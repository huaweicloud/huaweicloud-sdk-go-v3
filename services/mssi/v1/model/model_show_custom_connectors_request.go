package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomConnectorsRequest Request Object
type ShowCustomConnectorsRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset int32 `json:"offset"`

	// 每页显示的条目数量
	Limit int32 `json:"limit"`

	// CustomConnectors的名字
	Name *string `json:"name,omitempty"`
}

func (o ShowCustomConnectorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomConnectorsRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomConnectorsRequest", string(data)}, " ")
}
