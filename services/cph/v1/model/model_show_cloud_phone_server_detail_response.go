package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCloudPhoneServerDetailResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id"`

	// 云手机服务器信息
	Servers []interface{} `json:"servers"`

	// 云手机服务器带宽信息的结构体数组
	BandWidths []interface{} `json:"band_widths"`

	// 云手机服务器卷信息的结构体数组
	Volumes []interface{} `json:"volumes"`

	// 服务器扩展网卡绑定的安全组信息 系统定义网络的服务器，该字段返回为空列表
	SecurityGroups []string `json:"security_groups"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowCloudPhoneServerDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneServerDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneServerDetailResponse", string(data)}, " ")
}
