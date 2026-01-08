package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeriodResp **参数解释**： 指标周期，单位是秒。如想了解各个云服务的指标原始周期可以参考[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg) **取值范围**： 枚举值。 - 0代表立即触发，仅限事件场景使用。 - 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算。 - 300代表指标按5分钟聚合周期为一个数据点参与告警计算。 - 1200代表指标按20分钟聚合周期为一个数据点参与告警计算。 - 3600代表指标按1小时聚合周期为一个数据点参与告警计算。 - 14400代表指标按4小时聚合周期为一个数据点参与告警计算。 - 86400代表指标按1天聚合周期为一个数据点参与告警计算。
type PeriodResp struct {
	value int32
}

type PeriodRespEnum struct {
	E_0     PeriodResp
	E_1     PeriodResp
	E_300   PeriodResp
	E_1200  PeriodResp
	E_3600  PeriodResp
	E_14400 PeriodResp
	E_86400 PeriodResp
}

func GetPeriodRespEnum() PeriodRespEnum {
	return PeriodRespEnum{
		E_0: PeriodResp{
			value: 0,
		}, E_1: PeriodResp{
			value: 1,
		}, E_300: PeriodResp{
			value: 300,
		}, E_1200: PeriodResp{
			value: 1200,
		}, E_3600: PeriodResp{
			value: 3600,
		}, E_14400: PeriodResp{
			value: 14400,
		}, E_86400: PeriodResp{
			value: 86400,
		},
	}
}

func (c PeriodResp) Value() int32 {
	return c.value
}

func (c PeriodResp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodResp) UnmarshalJSON(b []byte) error {
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
