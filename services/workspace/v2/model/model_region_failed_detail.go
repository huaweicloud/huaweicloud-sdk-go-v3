package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionFailedDetail 区域操作失败详情。
type RegionFailedDetail struct {

	// 失败的区域标识。
	Region *string `json:"region,omitempty"`

	// 错误码，格式 WKS.XXXXXXXX。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o RegionFailedDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionFailedDetail struct{}"
	}

	return strings.Join([]string{"RegionFailedDetail", string(data)}, " ")
}
