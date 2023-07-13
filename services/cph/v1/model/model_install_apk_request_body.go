package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallApkRequestBody 安装应用请求体。
type InstallApkRequestBody struct {

	// 安装应用固定填写install。
	Command string `json:"command"`

	// 指定OBS桶中的apk文件（需要提前上传到指定桶中）。 只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）。推送的文件只支持apk文件类型。  单apk场景，只能传一个apk，最大长度为1024字节；多apk场景，最多传50个apk，中间用空格分开，最大长度8100字节。
	Content string `json:"content"`

	// 云手机ID列表。 server_ids参数不存在时必选，同时存在只处理phone_ids。
	PhoneIds *[]string `json:"phone_ids,omitempty"`

	// 云手机服务器ID列表。 phone_ids参数不存在时必选，同时存在只处理phone_ids。
	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o InstallApkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallApkRequestBody struct{}"
	}

	return strings.Join([]string{"InstallApkRequestBody", string(data)}, " ")
}
