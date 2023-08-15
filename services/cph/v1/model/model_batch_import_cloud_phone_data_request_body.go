package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportCloudPhoneDataRequestBody 导入云手机数据请求体
type BatchImportCloudPhoneDataRequestBody struct {

	// 待导入数据的云手机信息
	RestoreInfos []RestoreInfo `json:"restore_infos"`
}

func (o BatchImportCloudPhoneDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportCloudPhoneDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchImportCloudPhoneDataRequestBody", string(data)}, " ")
}
