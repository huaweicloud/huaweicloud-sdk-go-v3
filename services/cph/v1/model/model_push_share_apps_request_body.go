package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PushShareAppsRequestBody struct {

	// 合法的应用包名
	PackageName string `json:"package_name"`

	// 是否预装应用，1：预装，0：不预装；默认不预装
	PreInstallApp *int32 `json:"pre_install_app,omitempty"`

	// 合法的OBS桶名，3-63个字符，只能由小写字母、数字、中划线（-）和小数点组成
	BucketName string `json:"bucket_name"`

	// 合法的OBS对象key，最大长度1024字符。 推送的文件只支持tar文件类型。推送时，按tar文件解压后的文件目录结构推送到手机
	ObjectPath string `json:"object_path"`

	// 云手机服务器ID列表
	ServerIds []string `json:"server_ids"`
}

func (o PushShareAppsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushShareAppsRequestBody struct{}"
	}

	return strings.Join([]string{"PushShareAppsRequestBody", string(data)}, " ")
}
