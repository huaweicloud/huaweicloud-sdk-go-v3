package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbUpgradePrecheck 数据库检查结果
type DbUpgradePrecheck struct {

	// 检查是否通过，0代表通过，1代表未通过
	Result *int32 `json:"result,omitempty"`

	// 检查项
	ChecksPerformed *[]DbCheckDetail `json:"checks_performed,omitempty"`
}

func (o DbUpgradePrecheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbUpgradePrecheck struct{}"
	}

	return strings.Join([]string{"DbUpgradePrecheck", string(data)}, " ")
}
