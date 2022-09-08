package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCloudPhoneDetailResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id"`

	// 云手机名称，不超过65个字符
	PhoneName string `json:"phone_name"`

	// 云手机服务器ID，不超过32个字节
	ServerId string `json:"server_id"`

	// 云手机的唯一标识，不超过32个字节
	PhoneId string `json:"phone_id"`

	// 云手机镜像ID，不超过32个字节
	ImageId string `json:"image_id"`

	// 云手机是否开启VNC服务 - true：开启 - false：关闭
	VncEnable string `json:"vnc_enable"`

	// 云手机规格名称，不超过64个字节
	PhoneModelName string `json:"phone_model_name"`

	// 云手机状态 - 0：创建中 - 1：创建中 - 2：运行中 - 3：重置中 - 4：重启中 - 6：冻结 - 7：正在关机 - 8：已关机 - -5：重置失败 - -6：重启失败 - -7：手机异常 - -8：创建失败 - -9：关机失败
	Status int32 `json:"status"`

	// 云手机访问信息
	AccessInfos []interface{} `json:"access_infos"`

	// 云手机属性字符串，不超过2048个字节
	Property string `json:"property"`

	// 订单以及产品相关信息
	Metadata *interface{} `json:"metadata"`

	// 创建时间 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	CreateTime string `json:"create_time"`

	// 更新时间 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime     string `json:"update_time"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCloudPhoneDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneDetailResponse", string(data)}, " ")
}
