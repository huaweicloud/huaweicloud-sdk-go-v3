package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PushShareFilesRequestBody struct {

	// 合法的OBS桶名，3-63个字符，只能由小写字母、数字、中划线（-）和小数点组成。仅推送共享存储接口使用。
	BucketName string `json:"bucket_name"`

	// 合法的OBS对象路径，最大长度1024字符。仅支持大小写字母、数字、小数点（.）、斜线（/）、中划线（-）、下划线（_）、加号（+）、井号（#）、等号等这些字符。 推送的文件只支持tar文件类型。推送时，按tar文件解压后的文件目录结构推送到手机。当前只支持/data和/cache目录推送。仅推送共享存储接口使用。
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
