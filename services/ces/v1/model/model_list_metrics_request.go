package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListMetricsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 指标的维度，目前最大支持3个维度，从0开始；维度格式为dim.{i}=key,value，最大值为256。 例如：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d；各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim0 *string `json:"dim.0,omitempty" xml:"dim.0"`

	// 指标的维度，目前最大支持3个维度，从0开始；维度格式为dim.{i}=key,value，最大值为256。 例如：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d；各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim1 *string `json:"dim.1,omitempty" xml:"dim.1"`

	// 指标的维度，目前最大支持3个维度，从0开始；维度格式为dim.{i}=key,value，最大值为256。 例如：instance_id,6f3c6f91-4b24-4e1b-b7d1-a94ac1cb011d；各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Dim2 *string `json:"dim.2,omitempty" xml:"dim.2"`

	// 取值范围(0,1000]，默认值为1000。  用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 指标名称，例如弹性云服务器监控指标中的cpu_util；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	MetricName *string `json:"metric_name,omitempty" xml:"metric_name"`

	// 指标命名空间，格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32；各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 用于标识结果排序方法。  取值说明，默认为desc：  asc，升序 desc，降序
	Order *ListMetricsRequestOrder `json:"order,omitempty" xml:"order"`

	// 分页起始值，格式为：namespace.metric_name.key:value 例如：start=SYS.ECS.cpu_util.instance_id:d9112af5-6913-4f3b-bd0a-3f96711e004d
	Start *string `json:"start,omitempty" xml:"start"`
}

func (o ListMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListMetricsRequest", string(data)}, " ")
}

type ListMetricsRequestOrder struct {
	value string
}

type ListMetricsRequestOrderEnum struct {
	ASC  ListMetricsRequestOrder
	DESC ListMetricsRequestOrder
}

func GetListMetricsRequestOrderEnum() ListMetricsRequestOrderEnum {
	return ListMetricsRequestOrderEnum{
		ASC: ListMetricsRequestOrder{
			value: "asc",
		},
		DESC: ListMetricsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListMetricsRequestOrder) Value() string {
	return c.value
}

func (c ListMetricsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetricsRequestOrder) UnmarshalJSON(b []byte) error {
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
