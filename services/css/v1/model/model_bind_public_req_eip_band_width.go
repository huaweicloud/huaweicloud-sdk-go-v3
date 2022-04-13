package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 公网带宽信息。
type BindPublicReqEipBandWidth struct {
	// 带宽大小。

	Size int32 `json:"size"`
}

func (o BindPublicReqEipBandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicReqEipBandWidth struct{}"
	}

	return strings.Join([]string{"BindPublicReqEipBandWidth", string(data)}, " ")
}
