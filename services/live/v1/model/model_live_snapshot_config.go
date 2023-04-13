package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LiveSnapshotConfig struct {

	// 直播推流域名
	Domain string `json:"domain"`

	// 应用名称
	AppName string `json:"app_name"`

	// 回调鉴权密钥值  长度范围：[32-128]  若需要使用回调鉴权功能，请配置鉴权密钥，否则，留空即可。
	AuthKey *string `json:"auth_key,omitempty"`

	// 截图频率  取值范围：[5-3600]  单位：秒
	TimeInterval int32 `json:"time_interval"`

	// 在OBS桶存储截图的方式：  - 0：实时截图，以时间戳命名截图文件，保存所有截图文件到OBS桶。例：snapshot/{domain}/{app_name}/{stream_name}/{UnixTimestamp}.jpg  - 1：覆盖截图，只保存最新的截图文件，新的截图会覆盖原来的截图文件。例：snapshot/{domain}/{app_name}/{stream_name}.jpg
	ObjectWriteMode int32 `json:"object_write_mode"`

	ObsLocation *ObsFileAddr `json:"obs_location"`

	// 是否启用回调通知 - on：启用。 - off：不启用。
	CallBackEnable *LiveSnapshotConfigCallBackEnable `json:"call_back_enable,omitempty"`

	// 通知服务器地址，必须是合法的URL且携带协议，协议支持http和https。截图完成后直播服务会向此地址推送截图状态信息。
	CallBackUrl *string `json:"call_back_url,omitempty"`

	// 通知服务器区域。 - mainland_china：表示中国大陆区域 - outside_mainland_china：表示中国大陆以外区域  为空则按默认处理。
	CallBackArea *LiveSnapshotConfigCallBackArea `json:"call_back_area,omitempty"`

	// 是否开启ssl校验服务服务端证书，仅当回调url使用https协议时有效
	CallBackSslVerify *bool `json:"call_back_ssl_verify,omitempty"`

	// 客户端校验服务端的受信证书，call_back_ssl_verify为ture时必需
	CallBackSslCa *string `json:"call_back_ssl_ca,omitempty"`

	// 截图存储文件命名规则， 仅支持jpg格式 包含 - '{obs_object}' OBS存储路径，即obs_location.object的值  - '{domain}' 域名 - '{app}' 应用名 - '{stream}'  流名  其中实时截图模式下  - '{unix_time}'  时间戳，秒 - '{unix_time_milli}'  时间戳，毫秒 - '{fmt_time_utc}'   格式化UTC时间, 格式：YYYYMMDDhhmmss, 如20060102070405 - '{fmt_time_local}'  格式化本地时间, 格式：YYYYMMDDhhmmss，如20060102150405 必选一个时间类型模板
	ImageObjectFormat *string `json:"image_object_format,omitempty"`
}

func (o LiveSnapshotConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveSnapshotConfig struct{}"
	}

	return strings.Join([]string{"LiveSnapshotConfig", string(data)}, " ")
}

type LiveSnapshotConfigCallBackEnable struct {
	value string
}

type LiveSnapshotConfigCallBackEnableEnum struct {
	ON  LiveSnapshotConfigCallBackEnable
	OFF LiveSnapshotConfigCallBackEnable
}

func GetLiveSnapshotConfigCallBackEnableEnum() LiveSnapshotConfigCallBackEnableEnum {
	return LiveSnapshotConfigCallBackEnableEnum{
		ON: LiveSnapshotConfigCallBackEnable{
			value: "on",
		},
		OFF: LiveSnapshotConfigCallBackEnable{
			value: "off",
		},
	}
}

func (c LiveSnapshotConfigCallBackEnable) Value() string {
	return c.value
}

func (c LiveSnapshotConfigCallBackEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveSnapshotConfigCallBackEnable) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type LiveSnapshotConfigCallBackArea struct {
	value string
}

type LiveSnapshotConfigCallBackAreaEnum struct {
	MAINLAND_CHINA         LiveSnapshotConfigCallBackArea
	OUTSIDE_MAINLAND_CHINA LiveSnapshotConfigCallBackArea
}

func GetLiveSnapshotConfigCallBackAreaEnum() LiveSnapshotConfigCallBackAreaEnum {
	return LiveSnapshotConfigCallBackAreaEnum{
		MAINLAND_CHINA: LiveSnapshotConfigCallBackArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: LiveSnapshotConfigCallBackArea{
			value: "outside_mainland_china",
		},
	}
}

func (c LiveSnapshotConfigCallBackArea) Value() string {
	return c.value
}

func (c LiveSnapshotConfigCallBackArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveSnapshotConfigCallBackArea) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
