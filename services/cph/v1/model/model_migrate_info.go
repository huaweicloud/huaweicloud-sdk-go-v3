package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 待迁移数据的云手机信息。
type MigrateInfo struct {

	// 源云手机ID
	SourcePhoneId string `json:"source_phone_id"`

	// 目标云手机ID
	TargetPhoneId string `json:"target_phone_id"`

	// 是否迁移原手机的属性到目标手机。为\"true\"时迁移（忽略大小写），不填写或者其他值，则不迁移
	IsMigrateProperty *string `json:"is_migrate_property,omitempty"`
}

func (o MigrateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateInfo struct{}"
	}

	return strings.Join([]string{"MigrateInfo", string(data)}, " ")
}
