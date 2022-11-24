package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type DeleteClusterRequestBody struct {

	// 集群需要保留的快照数
	KeepLastManualSnapshot int32 `json:"keep_last_manual_snapshot"`
}

func (o DeleteClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequestBody", string(data)}, " ")
}
