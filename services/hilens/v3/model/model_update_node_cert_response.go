package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeCertResponse Response Object
type UpdateNodeCertResponse struct {

	// 设备ID
	Id *string `json:"id,omitempty"`

	// 设备名称，只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name *string `json:"name,omitempty"`

	// 将设备配置和证书文件node.conf/certificate/private_key打成.tar.gz包后用base64编码的字符串。node.conf包含节点信息配置
	Package        *string `json:"package,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNodeCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeCertResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeCertResponse", string(data)}, " ")
}
