package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchImportCloudPhoneDataRequestBody struct {

	// 待导入数据的云手机信息
	RestoreInfos []BatchImportCloudPhoneDataRequestBodyRestoreInfos `json:"restore_infos" xml:"restore_infos"`
}

func (o BatchImportCloudPhoneDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportCloudPhoneDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchImportCloudPhoneDataRequestBody", string(data)}, " ")
}
