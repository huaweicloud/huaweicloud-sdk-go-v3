package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchExportCloudPhoneDataRequestBody struct {

	// 待导出数据的云手机信息
	StorageInfos []BatchExportCloudPhoneDataRequestBodyStorageInfos `json:"storage_infos"`
}

func (o BatchExportCloudPhoneDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportCloudPhoneDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchExportCloudPhoneDataRequestBody", string(data)}, " ")
}
