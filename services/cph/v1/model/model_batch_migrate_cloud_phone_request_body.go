package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchMigrateCloudPhoneRequestBody struct {

	// 待迁移数据的云手机信息
	MigrateInfos []BatchMigrateCloudPhoneRequestBodyMigrateInfos `json:"migrate_infos"`
}

func (o BatchMigrateCloudPhoneRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateCloudPhoneRequestBody struct{}"
	}

	return strings.Join([]string{"BatchMigrateCloudPhoneRequestBody", string(data)}, " ")
}
