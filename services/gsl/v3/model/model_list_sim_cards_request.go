package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListSimCardsRequest struct {

	// 查询关键标识类型： 1.容器ID(不同类型卡含义如下:ICCID(实体卡)，EID（eSIM）CID（vSIM)) 2.批次号 3.设备IMEI
	MainSearchType *int32 `json:"main_search_type,omitempty"`

	// 查询关键标识值：根据查询关键标识类型进行查询，例如想根据ICCID=xxx进行查询，则main_search_type=1&main_search_key=xxx
	MainSearchKey *string `json:"main_search_key,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`

	// sim卡状态：  10.可测试  11.未激活  13.可激活  14.已停用  20.在用  30.已拆机
	SimStatus *int32 `json:"sim_status,omitempty"`

	// 设备状态: 1.注册 2.重启 3.在线 4.离线
	DeviceStatus *int32 `json:"device_status,omitempty"`

	// sim卡类型：  1.vSIM  2.eSIM  3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 排序的顺序，asc表示顺序排序，desc表示倒序排序，不传则默认asc
	Order *ListSimCardsRequestOrder `json:"order,omitempty"`

	// 排序的属性，目前支持:cid（容器ID）、flow_used（已用流量）、flow_left（剩余流量）、act_date（激活时间）、expire_time（到期时间）
	Sort *ListSimCardsRequestSort `json:"sort,omitempty"`

	// MSISDN
	Msisdn *string `json:"msisdn,omitempty"`

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

	// 最小使用流量(MB)
	MinUsedFlow *int64 `json:"min_used_flow,omitempty"`

	// 最大使用流量(MB)
	MaxUsedFlow *int64 `json:"max_used_flow,omitempty"`

	// 最小剩余流量(MB)
	MinLeftFlow *int64 `json:"min_left_flow,omitempty"`

	// 最大剩余流量(MB)
	MaxLeftFlow *int64 `json:"max_left_flow,omitempty"`

	// 是否已实名认证: true表示是，false表示否，系统SIM卡实名认证状态非实时。
	RealNamed *bool `json:"real_named,omitempty"`

	// 订单号
	OrderId *int64 `json:"order_id,omitempty"`

	// 是否过滤停机保号的卡
	FilterDowntimePeriod *bool `json:"filter_downtime_period,omitempty"`
}

func (o ListSimCardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimCardsRequest struct{}"
	}

	return strings.Join([]string{"ListSimCardsRequest", string(data)}, " ")
}

type ListSimCardsRequestOrder struct {
	value string
}

type ListSimCardsRequestOrderEnum struct {
	ASC  ListSimCardsRequestOrder
	DESC ListSimCardsRequestOrder
}

func GetListSimCardsRequestOrderEnum() ListSimCardsRequestOrderEnum {
	return ListSimCardsRequestOrderEnum{
		ASC: ListSimCardsRequestOrder{
			value: "asc",
		},
		DESC: ListSimCardsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListSimCardsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSimCardsRequestOrder) UnmarshalJSON(b []byte) error {
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

type ListSimCardsRequestSort struct {
	value string
}

type ListSimCardsRequestSortEnum struct {
	CID         ListSimCardsRequestSort
	FLOW_USED   ListSimCardsRequestSort
	FLOW_LEFT   ListSimCardsRequestSort
	ACT_DATE    ListSimCardsRequestSort
	EXPIRE_TIME ListSimCardsRequestSort
}

func GetListSimCardsRequestSortEnum() ListSimCardsRequestSortEnum {
	return ListSimCardsRequestSortEnum{
		CID: ListSimCardsRequestSort{
			value: "cid",
		},
		FLOW_USED: ListSimCardsRequestSort{
			value: "flow_used",
		},
		FLOW_LEFT: ListSimCardsRequestSort{
			value: "flow_left",
		},
		ACT_DATE: ListSimCardsRequestSort{
			value: "act_date",
		},
		EXPIRE_TIME: ListSimCardsRequestSort{
			value: "expire_time",
		},
	}
}

func (c ListSimCardsRequestSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSimCardsRequestSort) UnmarshalJSON(b []byte) error {
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
