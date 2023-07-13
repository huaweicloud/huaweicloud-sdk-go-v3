package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PushShareFilesRequestBody struct {

	// 合法的OBS桶名，3-63个字符，只能由小写字母、数字、中划线（-）和小数点组成。仅推送共享存储接口使用。
	BucketName string `json:"bucket_name"`

	// 合法的OBS对象key，最大长度1024字符。 推送的文件只支持tar文件类型。推送时，按tar文件解压后的文件目录结构推送到手机。当前只支持/data和/cache目录推送。仅推送共享存储接口使用。
	ObjectPath string `json:"object_path"`

	// 云手机服务器ID列表。
	ServerIds []string `json:"server_ids"`
}

func (o PushShareFilesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushShareFilesRequestBody struct{}"
	}

	return strings.Join([]string{"PushShareFilesRequestBody", string(data)}, " ")
}
