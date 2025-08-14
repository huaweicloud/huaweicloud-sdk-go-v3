package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeHdaVersionReq 批量升级HDA版本请求。
type BatchUpgradeHdaVersionReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]。
	Items []string `json:"items"`
}

func (o BatchUpgradeHdaVersionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeHdaVersionReq struct{}"
	}

	return strings.Join([]string{"BatchUpgradeHdaVersionReq", string(data)}, " ")
}
