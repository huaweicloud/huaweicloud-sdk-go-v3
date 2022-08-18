package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchImportCloudPhoneDataRequestBodyRestoreInfos struct {

	// 云手机ID
	PhoneId string `json:"phone_id"`

	// 存储云手机数据的OBS桶名。 合法的OBS桶名，3-63个字符，只能由小写字母、数字、中划线（-）和小数点（.）组成
	BucketName string `json:"bucket_name"`

	// 数据存储的OBS路径名。bucket_name与object_path长度累加要超过128。 合法的OBS对象key，最大长度1024字符
	ObjectPath string `json:"object_path"`
}

func (o BatchImportCloudPhoneDataRequestBodyRestoreInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportCloudPhoneDataRequestBodyRestoreInfos struct{}"
	}

	return strings.Join([]string{"BatchImportCloudPhoneDataRequestBodyRestoreInfos", string(data)}, " ")
}
