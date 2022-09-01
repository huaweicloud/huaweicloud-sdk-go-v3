package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchExportCloudPhoneDataRequestBodyStorageInfos struct {

	// 云手机ID
	PhoneId string `json:"phone_id" xml:"phone_id"`

	// 需要导出数据的存储路径 绝对路径，最大长度4096字节；目前只支持大小写字母、数字、小数点（.）、斜线（/）、中划线（-）、空格这些字符
	IncludeFiles []string `json:"include_files" xml:"include_files"`

	// 不能导出数据的存储路径。exclude_files优先级比include_files高，如果冲突，exclude_files生效。 参数可选，如果指定参数，则不能为空。 路径要求同include_files
	ExcludeFiles *[]string `json:"exclude_files,omitempty" xml:"exclude_files"`

	// 导出数据存储的OBS桶名。 合法的OBS桶名，3-63个字符，只能由小写字母、数字、中划线（-）和小数点（.）组成
	BucketName string `json:"bucket_name" xml:"bucket_name"`

	// 导出数据存储的OBS路径名。bucket_name与object_path长度累加要超过128。 符合OBS的路径名规范，最大长度1024字符
	ObjectPath string `json:"object_path" xml:"object_path"`
}

func (o BatchExportCloudPhoneDataRequestBodyStorageInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportCloudPhoneDataRequestBodyStorageInfos struct{}"
	}

	return strings.Join([]string{"BatchExportCloudPhoneDataRequestBodyStorageInfos", string(data)}, " ")
}
