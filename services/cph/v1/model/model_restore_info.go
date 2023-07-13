package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreInfo 待重建数据的云手机信息。
type RestoreInfo struct {

	// 云手机ID
	PhoneId string `json:"phone_id"`

	// 导出数据存储的OBS桶名。 合法的OBS桶名，3-63个字符，只能由小写字母、数字、中划线（-）和小数点（.）组成
	BucketName string `json:"bucket_name"`

	// 导出数据存储的OBS路径名。 符合OBS的路径名规范，最大长度1024字符
	ObjectPath string `json:"object_path"`
}

func (o RestoreInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInfo struct{}"
	}

	return strings.Join([]string{"RestoreInfo", string(data)}, " ")
}
