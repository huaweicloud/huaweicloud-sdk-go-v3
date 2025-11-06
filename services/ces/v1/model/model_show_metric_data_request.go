package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMetricDataRequest Request Object
type ShowMetricDataRequest struct {

	// 指标命名空间，如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。
	Namespace string `json:"namespace"`

	// 资源的监控指标名称，如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	MetricName string `json:"metric_name"`

	// 指标的第一层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.0=key,value，如mongodb_cluster_id,4270ff17-aba3-4138-89fa-820594c39755；key为指标的维度信息，如：文档数据库服务，则第一层维度为mongodb_cluster_id，value为文档数据库实例ID；各服务资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。
	Dim0 string `json:"dim.0"`

	// 指标的第二层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.1=key,value，如mongos_instance_id,c65d39d7-185c-4616-9aca-ad65703b15f9；key为指标的维度信息，如：文档数据库服务，则第二层维度为mongos_instance_id，value为文档数据库集群实例下的mongos节点ID；各资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。
	Dim1 *string `json:"dim.1,omitempty"`

	// 指标的第三层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.2=key,value，如mongod_primary_instance_id,5f9498e9-36f8-4317-9ea1-ebe28cba99b4；key为指标的维度信息，如：文档数据库服务，则第三层维度为mongod_primary_instance_id，value为文档数据库实例下的主节点ID；各资源的指标维度名称可查看：“[服务指标维度]ces_03_0059.xml)”。
	Dim2 *string `json:"dim.2,omitempty"`

	// 指标的第四层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.3=key,value，如mongod_secondary_instance_id,b46fa2c7-aac6-4ae3-9337-f4ea97f885cb；key为指标的维度信息，如：文档数据库服务，则第四层维度为mongod_secondary_instance_id，value为文档数据库实例下的备节点ID；各资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。
	Dim3 *string `json:"dim.3,omitempty"`

	// 聚合方式。average：平均值，variance：方差，min：最小值，max：最大值，sum：求和。
	Filter ShowMetricDataRequestFilter `json:"filter"`

	// 指标监控数据的聚合粒度，取值范围：1，60，300，1200，3600，14400，86400；1为监控资源的实时数据；60为聚合1分钟粒度数据，表示1分钟一个数据点；300为聚合5分钟粒度数据，表示5分钟一个数据点；1200为聚合20分钟粒度数据，表示20分钟一个数据点；3600为聚合1小时粒度数据，表示1小时一个数据点；14400为聚合4小时粒度数据，表示4小时一个数据点；86400为聚合1天粒度数据，表示1天一个数据点；聚合解释可查看：“[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)”。
	Period ShowMetricDataRequestPeriod `json:"period"`

	// 查询数据起始时间，UNIX时间戳，单位毫秒。建议from的值相对于当前时间向前偏移至少1个周期。由于聚合运算的过程是将一个聚合周期范围内的数据点聚合到周期起始边界上，如果将from和to的范围设置在聚合周期内，会因为聚合未完成而造成查询数据为空，所以建议from参数相对于当前时间向前偏移至少1个周期。以5分钟聚合周期为例：假设当前时间点为10:35，10:30~10:35之间的原始数据会被聚合到10:30这个点上，所以查询5分钟数据点时from参数应为10:30或之前。云监控会根据所选择的聚合粒度向前取整from参数；如：1607146998177。
	From int64 `json:"from"`

	// 查询数据截止时间UNIX时间戳，单位毫秒。from必须小于to；如：1607150598177。
	To int64 `json:"to"`
}

func (o ShowMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricDataRequest", string(data)}, " ")
}

type ShowMetricDataRequestFilter struct {
	value string
}

type ShowMetricDataRequestFilterEnum struct {
	AVERAGE  ShowMetricDataRequestFilter
	VARIANCE ShowMetricDataRequestFilter
	MIN      ShowMetricDataRequestFilter
	MAX      ShowMetricDataRequestFilter
	SUM      ShowMetricDataRequestFilter
}

func GetShowMetricDataRequestFilterEnum() ShowMetricDataRequestFilterEnum {
	return ShowMetricDataRequestFilterEnum{
		AVERAGE: ShowMetricDataRequestFilter{
			value: "average",
		},
		VARIANCE: ShowMetricDataRequestFilter{
			value: "variance",
		},
		MIN: ShowMetricDataRequestFilter{
			value: "min",
		},
		MAX: ShowMetricDataRequestFilter{
			value: "max",
		},
		SUM: ShowMetricDataRequestFilter{
			value: "sum",
		},
	}
}

func (c ShowMetricDataRequestFilter) Value() string {
	return c.value
}

func (c ShowMetricDataRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricDataRequestFilter) UnmarshalJSON(b []byte) error {
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

type ShowMetricDataRequestPeriod struct {
	value int32
}

type ShowMetricDataRequestPeriodEnum struct {
	E_1     ShowMetricDataRequestPeriod
	E_60    ShowMetricDataRequestPeriod
	E_300   ShowMetricDataRequestPeriod
	E_1200  ShowMetricDataRequestPeriod
	E_3600  ShowMetricDataRequestPeriod
	E_14400 ShowMetricDataRequestPeriod
	E_86400 ShowMetricDataRequestPeriod
}

func GetShowMetricDataRequestPeriodEnum() ShowMetricDataRequestPeriodEnum {
	return ShowMetricDataRequestPeriodEnum{
		E_1: ShowMetricDataRequestPeriod{
			value: 1,
		}, E_60: ShowMetricDataRequestPeriod{
			value: 60,
		}, E_300: ShowMetricDataRequestPeriod{
			value: 300,
		}, E_1200: ShowMetricDataRequestPeriod{
			value: 1200,
		}, E_3600: ShowMetricDataRequestPeriod{
			value: 3600,
		}, E_14400: ShowMetricDataRequestPeriod{
			value: 14400,
		}, E_86400: ShowMetricDataRequestPeriod{
			value: 86400,
		},
	}
}

func (c ShowMetricDataRequestPeriod) Value() int32 {
	return c.value
}

func (c ShowMetricDataRequestPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricDataRequestPeriod) UnmarshalJSON(b []byte) error {
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
