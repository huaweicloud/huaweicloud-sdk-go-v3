package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchPeriod **参数解释** 监控数据的聚合粒度，聚合解释可查看：“[[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)](tag:hc)[[聚合含义](https://support.huaweicloud.com/intl/zh-cn/ces_faq/ces_faq_0009.html)](tag:hk)[[聚合含义](https://support.huaweicloud.com/eu/ces_faq/ces_faq_0009.html)](tag:hws_eu)[[聚合含义](https://docs.otc.t-systems.com/usermanual/ces/ces_faq_0009.html)](tag:dt,dt_test)[《云监控服务用户指南》中“什么是聚合？”章节](tag:ax,cmcc,ctc,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)”。 **约束限制** 不涉及 **取值范围** 枚举值： - 1 监控资源的实时数据 - 60 聚合1分钟粒度数据，表示1分钟一个数据点 - 300 聚合5分钟粒度数据，表示5分钟一个数据点 - 1200 聚合20分钟粒度数据，表示20分钟一个数据点 - 3600 聚合1小时粒度数据，表示1小时一个数据点 - 14400 聚合4小时粒度数据，表示4小时一个数据点 - 86400 聚合1天粒度数据，表示1天一个数据点 **默认取值** 不涉及
type BatchPeriod struct {
	value string
}

type BatchPeriodEnum struct {
	E_1     BatchPeriod
	E_60    BatchPeriod
	E_300   BatchPeriod
	E_1200  BatchPeriod
	E_3600  BatchPeriod
	E_14400 BatchPeriod
	E_86400 BatchPeriod
}

func GetBatchPeriodEnum() BatchPeriodEnum {
	return BatchPeriodEnum{
		E_1: BatchPeriod{
			value: "1",
		},
		E_60: BatchPeriod{
			value: "60",
		},
		E_300: BatchPeriod{
			value: "300",
		},
		E_1200: BatchPeriod{
			value: "1200",
		},
		E_3600: BatchPeriod{
			value: "3600",
		},
		E_14400: BatchPeriod{
			value: "14400",
		},
		E_86400: BatchPeriod{
			value: "86400",
		},
	}
}

func (c BatchPeriod) Value() string {
	return c.value
}

func (c BatchPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchPeriod) UnmarshalJSON(b []byte) error {
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
