package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BatchListMetricDataRequestBody struct {
	// 指标数据。数组长度最大500

	Metrics []MetricInfo `json:"metrics"`
	// 指标监控数据的聚合粒度，取值范围：1，300，1200，3600，14400，86400；1为监控资源的实时数据；300为聚合5分钟粒度数据，表示5分钟一个数据点；1200为聚合20分钟粒度数据，表示20分钟一个数据点；3600为聚合1小时粒度数据，表示1小时一个数据点；14400为聚合4小时粒度数据，表示4小时一个数据点；86400为聚合1天粒度数据，表示1天一个数据点；聚合解释可查看：“[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)”。

	Period string `json:"period"`
	// 数据聚合方式。  支持的值为max, min, average, sum, variance；max为最大值，min为最小值，average为平均值，sum为和，variance为方差值。

	Filter string `json:"filter"`
	// 查询数据起始时间，UNIX时间戳，单位毫秒。建议from的值相对于当前时间向前偏移至少1个周期。由于聚合运算的过程是将一个聚合周期范围内的数据点聚合到周期起始边界上，如果将from和to的范围设置在聚合周期内，会因为聚合未完成而造成查询数据为空，所以建议from参数相对于当前时间向前偏移至少1个周期。以5分钟聚合周期为例：假设当前时间点为10:35，10:30~10:35之间的原始数据会被聚合到10:30这个点上，所以查询5分钟数据点时from参数应为10:30或之前。 说明： 云监控会根据所选择的聚合粒度向前取整from参数；如：1607146998177

	From int64 `json:"from"`
	// 查询数据截止时间UNIX时间戳，单位毫秒。from必须小于to；如：1607150598177。

	To int64 `json:"to"`
}

func (o BatchListMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequestBody", string(data)}, " ")
}
