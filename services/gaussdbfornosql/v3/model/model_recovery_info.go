package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecoveryInfo 参数解释： '备份文件所在OBS信息。' 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
type RecoveryInfo struct {

	// 参数解释： '备份文件所在OBS bucket。' 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
	BucketName string `json:"bucket_name"`

	// 参数解释： '备份文件名。' 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
	Files []string `json:"files"`
}

func (o RecoveryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecoveryInfo struct{}"
	}

	return strings.Join([]string{"RecoveryInfo", string(data)}, " ")
}
