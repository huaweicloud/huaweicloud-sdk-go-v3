package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchMigrateCloudPhoneRequestBodyMigrateInfos struct {

	// 源云手机ID
	SourcePhoneId string `json:"source_phone_id"`

	// 目标云手机ID
	TargetPhoneId string `json:"target_phone_id"`

	// 是否迁移原手机的属性到目标手机。为\"true\"时迁移（忽略大小写），不填写或者填写为其他值，则不迁移
	IsMigrateProperty *string `json:"is_migrate_property,omitempty"`
}

func (o BatchMigrateCloudPhoneRequestBodyMigrateInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateCloudPhoneRequestBodyMigrateInfos struct{}"
	}

	return strings.Join([]string{"BatchMigrateCloudPhoneRequestBodyMigrateInfos", string(data)}, " ")
}
