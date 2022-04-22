package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowMetricDataRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type"`

	// 指标命名空间，如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Namespace string `json:"namespace"`

	// 资源的监控指标名称，如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	MetricName string `json:"metric_name"`

	// 指标的第一层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.0=key,value，如dim.0=mongodb_cluster_id,4270ff17-aba3-4138-89fa-820594c39755；key为指标的维度信息，如：文档数据库服务，则第一层维度为mongodb_cluster_id，value为文档数据库实例ID；各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim0 string `json:"dim.0"`

	// 指标的第二层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.1=key,value，如dim.1=mongos_instance_id,c65d39d7-185c-4616-9aca-ad65703b15f9；key为指标的维度信息，如：文档数据库服务，则第二层维度为mongos_instance_id，value为文档数据库集群实例下的mongos节点ID；各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim1 *string `json:"dim.1,omitempty"`

	// 指标的第三层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.2=key,value，如dim.2=mongod_primary_instance_id,5f9498e9-36f8-4317-9ea1-ebe28cba99b4；key为指标的维度信息，如：文档数据库服务，则第三层维度为mongod_primary_instance_id，value为文档数据库实例下的主节点ID；各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim2 *string `json:"dim.2,omitempty"`

	// 指标的第四层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.3=key,value，如dim.3=mongod_secondary_instance_id,b46fa2c7-aac6-4ae3-9337-f4ea97f885cb；key为指标的维度信息，如：文档数据库服务，则第四层维度为mongod_secondary_instance_id，value为文档数据库实例下的备节点ID；各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim3 *string `json:"dim.3,omitempty"`

	// 数据聚合方式。支持的值为max, min, average, sum, variance。
	Filter ShowMetricDataRequestFilter `json:"filter"`

	// 指标监控数据的聚合粒度，取值范围：1，300，1200，3600，14400，86400；1为监控资源的实时数据；300为聚合5分钟粒度数据，表示5分钟一个数据点；1200为聚合20分钟粒度数据，表示20分钟一个数据点；3600为聚合1小时粒度数据，表示1小时一个数据点；14400为聚合4小时粒度数据，表示4小时一个数据点；86400为聚合1天粒度数据，表示1天一个数据点；聚合解释可查看：“[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)”。
	Period int32 `json:"period"`

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
	MAX      ShowMetricDataRequestFilter
	MIN      ShowMetricDataRequestFilter
	AVERAGE  ShowMetricDataRequestFilter
	SUM      ShowMetricDataRequestFilter
	VARIANCE ShowMetricDataRequestFilter
}

func GetShowMetricDataRequestFilterEnum() ShowMetricDataRequestFilterEnum {
	return ShowMetricDataRequestFilterEnum{
		MAX: ShowMetricDataRequestFilter{
			value: "max",
		},
		MIN: ShowMetricDataRequestFilter{
			value: "min",
		},
		AVERAGE: ShowMetricDataRequestFilter{
			value: "average",
		},
		SUM: ShowMetricDataRequestFilter{
			value: "sum",
		},
		VARIANCE: ShowMetricDataRequestFilter{
			value: "variance",
		},
	}
}

func (c ShowMetricDataRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMetricDataRequestFilter) UnmarshalJSON(b []byte) error {
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
