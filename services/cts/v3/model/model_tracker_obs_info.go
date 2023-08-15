package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TrackerObsInfo 转储桶信息。
type TrackerObsInfo struct {

	// 标识OBS桶名称。由数字或字母开头，支持小写字母、数字、“-”、“.”，长度为3～63个字符。
	BucketName *string `json:"bucket_name,omitempty"`

	// 标识需要存储于OBS的日志文件前缀，0-9，a-z，A-Z，'-'，'.'，'_'长度为0～64字符。
	FilePrefixName *string `json:"file_prefix_name,omitempty"`

	// 是否支持新建OBS桶。   值为“true”时，表示新创建OBS桶存储事件文件；   值为“false”时，选择已存在的OBS桶存储事件文件。
	IsObsCreated *bool `json:"is_obs_created,omitempty"`

	// 标识配置桶内对象存储周期。 当\"tracker_type\"参数值为\"data\"时该参数值有效。
	BucketLifecycle *int32 `json:"bucket_lifecycle,omitempty"`

	// 压缩类型。包括不压缩（json），压缩（gzip）两种状态。默认为gzip格式。
	CompressType *TrackerObsInfoCompressType `json:"compress_type,omitempty"`

	// 路径按云服务划分，打开后转储文件路径中将增加云服务名。默认为true。
	IsSortByService *bool `json:"is_sort_by_service,omitempty"`
}

func (o TrackerObsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerObsInfo struct{}"
	}

	return strings.Join([]string{"TrackerObsInfo", string(data)}, " ")
}

type TrackerObsInfoCompressType struct {
	value string
}

type TrackerObsInfoCompressTypeEnum struct {
	GZIP TrackerObsInfoCompressType
	JSON TrackerObsInfoCompressType
}

func GetTrackerObsInfoCompressTypeEnum() TrackerObsInfoCompressTypeEnum {
	return TrackerObsInfoCompressTypeEnum{
		GZIP: TrackerObsInfoCompressType{
			value: "gzip",
		},
		JSON: TrackerObsInfoCompressType{
			value: "json",
		},
	}
}

func (c TrackerObsInfoCompressType) Value() string {
	return c.value
}

func (c TrackerObsInfoCompressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrackerObsInfoCompressType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
