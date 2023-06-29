package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TransferDetail 日志转储详细信息
type TransferDetail struct {

	// OBS转储时间
	ObsPeriod TransferDetailObsPeriod `json:"obs_period"`

	// OBS转储KMS秘钥ID。若OBS转储未加密则不返回此字段
	ObsEncryptedId *string `json:"obs_encrypted_id,omitempty"`

	// OBS转储日志文件前缀
	ObsPrefixName *string `json:"obs_prefix_name,omitempty"`

	// OBS转储单位
	ObsPeriodUnit string `json:"obs_period_unit"`

	// OBS转储路径，指OBS日志桶中的路径
	ObsTransferPath *string `json:"obs_transfer_path,omitempty"`

	// OBS企业项目ID
	ObsEpsId *string `json:"obs_eps_id,omitempty"`

	// OBS日志桶名称
	ObsBucketName string `json:"obs_bucket_name"`

	// OBS是否开启加密。
	ObsEncryptedEnable *bool `json:"obs_encrypted_enable,omitempty"`

	// OBS转储自定义转储路径
	ObsDirPreFixName *string `json:"obs_dir_pre_fix_name,omitempty"`

	// DIS转储通道ID
	DisId *string `json:"dis_id,omitempty"`

	// DIS转储通道名称
	DisName *string `json:"dis_name,omitempty"`

	// DMS转储kafka ID
	KafkaId *string `json:"kafka_id,omitempty"`

	// DMS转储kafka topic
	KafkaTopic *string `json:"kafka_topic,omitempty"`

	// OBS转储时区。如果选择该参数，则必须选择obs_time_zone_id。
	ObsTimeZone *string `json:"obs_time_zone,omitempty"`

	// OBS转储时区ID。参数选择参考OBS转储时区表。如果选择该参数，则必须选择obs_time_zone。
	ObsTimeZoneId *string `json:"obs_time_zone_id,omitempty"`

	// 若开启tag投递，该字段必须包含主机信息：hostIP、hostId、hostName、pathFile、collectTime；  公共字段有：logStreamName、regionName、logGroupName、projectId，为可选填；  开启转储标签：streamTag，可选填
	Tags *[]string `json:"tags,omitempty"`
}

func (o TransferDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferDetail struct{}"
	}

	return strings.Join([]string{"TransferDetail", string(data)}, " ")
}

type TransferDetailObsPeriod struct {
	value int32
}

type TransferDetailObsPeriodEnum struct {
	E_1  TransferDetailObsPeriod
	E_2  TransferDetailObsPeriod
	E_3  TransferDetailObsPeriod
	E_5  TransferDetailObsPeriod
	E_6  TransferDetailObsPeriod
	E_12 TransferDetailObsPeriod
	E_30 TransferDetailObsPeriod
}

func GetTransferDetailObsPeriodEnum() TransferDetailObsPeriodEnum {
	return TransferDetailObsPeriodEnum{
		E_1: TransferDetailObsPeriod{
			value: 1,
		}, E_2: TransferDetailObsPeriod{
			value: 2,
		}, E_3: TransferDetailObsPeriod{
			value: 3,
		}, E_5: TransferDetailObsPeriod{
			value: 5,
		}, E_6: TransferDetailObsPeriod{
			value: 6,
		}, E_12: TransferDetailObsPeriod{
			value: 12,
		}, E_30: TransferDetailObsPeriod{
			value: 30,
		},
	}
}

func (c TransferDetailObsPeriod) Value() int32 {
	return c.value
}

func (c TransferDetailObsPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransferDetailObsPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
