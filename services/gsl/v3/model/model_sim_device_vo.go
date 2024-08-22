package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimDeviceVo struct {

	// sim卡id
	SimCardId *int64 `json:"sim_card_id,omitempty"`

	// 账户id
	AccountId *string `json:"account_id,omitempty"`

	// 容器ID:不同类型卡含义如下 iccid(实体卡)，eid（eSIM）cid（vSIM)
	Cid *string `json:"cid,omitempty"`

	// 流量池ID
	SimPoolId *int64 `json:"sim_pool_id,omitempty"`

	// 设备IMEI
	Imei *string `json:"imei,omitempty"`

	// sim卡状态：  10.可测试  11.未激活  13.可激活  14.已停用  20.在用  30.已拆机
	SimStatus *int32 `json:"sim_status,omitempty"`

	// 设备状态：1.注册 2.重启 3.在线 4.离线 (该参数只有ESIM、VSIM返回, 实体卡返回null)
	DeviceStatus *int32 `json:"device_status,omitempty"`

	// 设备模组 (该参数只有ESIM、VSIM返回, 实体卡返回null)
	DeviceModel *string `json:"device_model,omitempty"`

	// 激活日期 例如2020-01-31T16:00:00.000Z
	ActDate *sdktime.SdkTime `json:"act_date,omitempty"`

	// 设备状态变更时间 例如2020-01-31T16:00:00.000Z (该参数只有ESIM、VSIM返回, 实体卡返回null)
	DeviceStatusDate *sdktime.SdkTime `json:"device_status_date,omitempty"`

	// 设备标识
	NodeId *string `json:"node_id,omitempty"`

	// 码号iccid
	Iccid *string `json:"iccid,omitempty"`

	// 网络类型
	NetworkType *string `json:"network_type,omitempty"`

	// 信号强度 (该参数只有ESIM、VSIM返回, 实体卡返回null)
	Dbm *string `json:"dbm,omitempty"`

	// 信号等级:1.差  2.良  3.良  4.优 (该参数只有ESIM、VSIM返回, 实体卡返回null)
	SignalLevel *string `json:"signal_level,omitempty"`

	// sim卡类型 1.vSIM  2.eSIM  3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 标签名
	TagNames *string `json:"tag_names,omitempty"`

	// 批次号
	OrderId *int64 `json:"order_id,omitempty"`

	// 到期时间 例如2021-06-30T00:00:00.000Z
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`

	// 在用套餐名
	PricePlanName *string `json:"price_plan_name,omitempty"`

	// 套餐订购实例ID
	SimPricePlanId *int64 `json:"sim_price_plan_id,omitempty"`

	// 剩余流量(单位M)，数据默认截止到昨日24点。
	FlowLeft *float64 `json:"flow_left,omitempty"`

	// 已用流量(单位M)，数据默认截止到昨日24点。
	FlowUsed *float64 `json:"flow_used,omitempty"`

	// 运营商状态 -1.正常（非停机状态） 1.停机（超流量停机） 2.停机（超流量阈值停机） 3.停机（流量池停机） 4.停机（套餐到期停机） 5.停机（主动停机） 6.停机（违规停机） 7.停机（机卡分离停机）
	OperatorStatus *int32 `json:"operator_status,omitempty"`

	// MSISDN
	Msisdn *string `json:"msisdn,omitempty"`

	// IMSI
	Imsi *string `json:"imsi,omitempty"`

	// 自定义属性一
	CustomerAttribute1 *string `json:"customer_attribute1,omitempty"`

	// 自定义属性二
	CustomerAttribute2 *string `json:"customer_attribute2,omitempty"`

	// 自定义属性三
	CustomerAttribute3 *string `json:"customer_attribute3,omitempty"`

	// 自定义属性四
	CustomerAttribute4 *string `json:"customer_attribute4,omitempty"`

	// 自定义属性五
	CustomerAttribute5 *string `json:"customer_attribute5,omitempty"`

	// 自定义属性六
	CustomerAttribute6 *string `json:"customer_attribute6,omitempty"`

	// 是否已实名认证: true表示是，false表示否，系统SIM卡实名认证状态非实时。
	RealNamed *bool `json:"real_named,omitempty"`

	// 是否单独断网 true:断网，false:未断网 （当前只支持联通、移动的组池卡，电信卡不限制）
	CutNetFlag *bool `json:"cut_net_flag,omitempty"`

	// 是否达量断网 true:达量断网，false:未达量断网 （当前只支持联通、移动的组池卡，电信卡不限制）
	ExceedCutNetFlag *bool `json:"exceed_cut_net_flag,omitempty"`

	// 达量断网阈值（单位MB 当前仅电信卡支持）
	ExceedCutNetQuota *int32 `json:"exceed_cut_net_quota,omitempty"`

	// 本月机卡绑定剩余次数（当前仅电信卡支持）
	ImeiBindRemainTimes *int32 `json:"imei_bind_remain_times,omitempty"`

	// 网络限制速率（单位Kbps,当前电信联通卡支持）
	SpeedValue *int32 `json:"speed_value,omitempty"`
}

func (o SimDeviceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimDeviceVo struct{}"
	}

	return strings.Join([]string{"SimDeviceVo", string(data)}, " ")
}
