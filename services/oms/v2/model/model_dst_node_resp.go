package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DstNodeResp struct {
	// 目的端桶的名称。

	Bucket *string `json:"bucket,omitempty"`
	// 目的端桶所处的区域。  请与Endpoint对应的区域保持一致。

	Region *string `json:"region,omitempty"`
}

func (o DstNodeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DstNodeResp struct{}"
	}

	return strings.Join([]string{"DstNodeResp", string(data)}, " ")
}
