package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceResp 提交升级任务id
type UpgradeInstanceResp struct {
}

func (o UpgradeInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceResp struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceResp", string(data)}, " ")
}
