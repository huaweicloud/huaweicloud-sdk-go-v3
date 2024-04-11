package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudPhoneDetailResponse Response Object
type ShowCloudPhoneDetailResponse struct {

	// 请求的唯一标识ID，不超过32个字节。
	RequestId *string `json:"request_id,omitempty"`

	// 云手机名称，不超过65个字符。
	PhoneName *string `json:"phone_name,omitempty"`

	// 云手机服务器ID，不超过32个字节。
	ServerId *string `json:"server_id,omitempty"`

	// 云手机的唯一标识，不超过32个字节。
	PhoneId *string `json:"phone_id,omitempty"`

	// 云手机镜像ID，不超过32个字节。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像版本。
	ImageVersion *string `json:"image_version,omitempty"`

	// 云手机是否开启VNC服务（过期） - true：开启 - false：关闭
	VncEnable *string `json:"vnc_enable,omitempty"`

	// 云手机规格名称，不超过64个字节。
	PhoneModelName *string `json:"phone_model_name,omitempty"`

	// 云手机状态。 - 1：创建中 - 2：运行中 - 3：重置中 - 4：重启中 - 6：冻结 - 7：正在关机 - 8：已关机 - -5：重置失败 - -6：重启失败 - -7：手机异常 - -8：创建失败 - -9：关机失败
	Status *int32 `json:"status,omitempty"`

	// 云手机访问信息。
	AccessInfos *[]PhoneAccessInfo `json:"access_infos,omitempty"`

	// 云手机属性字符串，不超过2048个字节。
	Property *string `json:"property,omitempty"`

	Metadata *ShowCloudPhoneDetailResponseBodyMetadata `json:"metadata,omitempty"`

	PhoneDataVolume *PhoneDataVolume `json:"phone_data_volume,omitempty"`

	// imei码。
	Imei *string `json:"imei,omitempty"`

	// 手机路由类型。 - direct：默认路由 - routing：路由到编码容器
	TrafficType *string `json:"traffic_type,omitempty"`

	// 手机物理磁盘是否独立。 - 0：不独立 - 1：独立
	VolumeMode *int32 `json:"volume_mode,omitempty"`

	// 云手机服务器所在的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 创建时间， 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间， 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCloudPhoneDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneDetailResponse", string(data)}, " ")
}
