package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunShellCommandRequestBody struct {

	// 手机管理命令 - 推送文件场景固定填写push - 安装apk场景固定填写install - 安装多apk场景固定填写install-multiple - 卸载apk场景固定填写uninstall - 执行命令固定写shell
	Command string `json:"command"`

	// - 推送文件场景：指定OBS桶中的tar文件。最大长度为1024字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）。推送的文件只支持tar文件类型。文件格式：obs://obs-bucket-name/obs-file-path/file.tar  - 安装apk场景：指定OBS桶中的apk文件（需要提前上传到指定桶中）。单apk场景，只能传一个apk。最大长度为1024字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）。推送的文件只支持apk文件类型。文件格式：obs://obs-bucket-name/obs-file-path/file.apk  - 安装多apk场景：指定OBS桶中的apk文件（需要提前上传到指定桶中）。多apk场景，最多传50个apk，中间用空格分开，最大长度8100字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）。推送的文件只支持apk文件类型。文件格式：obs://obs-bucket-name/obs-file-path/file.apk  - 卸载apk场景：待卸载的APP名称。最大长度为1024字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）  - 执行命令场景：待执行的命令。最大长度为1024字节，只支持大小写字母、数字、下划线（_）、点（.）、斜线（/）、冒号（:）、中划线（-）
	Content string `json:"content"`

	// 云手机ID列表 server_ids参数不存在时必选
	PhoneIds *[]string `json:"phone_ids,omitempty"`

	// 云手机服务器ID列表 phone_ids参数不存在时必选
	ServerIds *[]string `json:"server_ids,omitempty"`
}

func (o RunShellCommandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunShellCommandRequestBody struct{}"
	}

	return strings.Join([]string{"RunShellCommandRequestBody", string(data)}, " ")
}
