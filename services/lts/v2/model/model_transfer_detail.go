package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 日志转储详细信息
type TransferDetail struct {

	// OBS转储时间
	ObsPeriod TransferDetailObsPeriod `json:"obs_period" xml:"obs_period"`

	// OBS转储KMS秘钥ID。若OBS转储未加密则不返回此字段
	ObsEncryptedId *string `json:"obs_encrypted_id,omitempty" xml:"obs_encrypted_id"`

	// OBS转储日志文件前缀
	ObsPrefixName *string `json:"obs_prefix_name,omitempty" xml:"obs_prefix_name"`

	// OBS转储单位
	ObsPeriodUnit string `json:"obs_period_unit" xml:"obs_period_unit"`

	// OBS转储路径，指OBS日志桶中的路径
	ObsTransferPath *string `json:"obs_transfer_path,omitempty" xml:"obs_transfer_path"`

	// OBS企业项目ID
	ObsEpsId *string `json:"obs_eps_id,omitempty" xml:"obs_eps_id"`

	// OBS日志桶名称
	ObsBucketName string `json:"obs_bucket_name" xml:"obs_bucket_name"`

	// OBS是否开启加密。
	ObsEncryptedEnable *bool `json:"obs_encrypted_enable,omitempty" xml:"obs_encrypted_enable"`

	// OBS转储自定义转储路径
	ObsDirPreFixName *string `json:"obs_dir_pre_fix_name,omitempty" xml:"obs_dir_pre_fix_name"`

	// DIS转储通道ID
	DisId *string `json:"dis_id,omitempty" xml:"dis_id"`

	// DIS转储通道名称
	DisName *string `json:"dis_name,omitempty" xml:"dis_name"`

	// DMS转储kafka ID
	KafkaId *string `json:"kafka_id,omitempty" xml:"kafka_id"`

	// DMS转储kafka topic
	KafkaTopic *string `json:"kafka_topic,omitempty" xml:"kafka_topic"`

	// OBS转储时区。如果选择该参数，则必须选择obs_time_zone_id。
	ObsTimeZone *string `json:"obs_time_zone,omitempty" xml:"obs_time_zone"`

	// OBS转储时区ID。参数选择参考OBS转储时区表。如果选择该参数，则必须选择obs_time_zone。
	ObsTimeZoneId *string `json:"obs_time_zone_id,omitempty" xml:"obs_time_zone_id"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
