package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypeUsageQuantity struct {

	// 总文件系统数量
	Total *int32 `json:"total,omitempty"`

	// 文件系统已用数量
	Usage *int32 `json:"usage,omitempty"`
}

func (o ShareTypeUsageQuantity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypeUsageQuantity struct{}"
	}

	return strings.Join([]string{"ShareTypeUsageQuantity", string(data)}, " ")
}
