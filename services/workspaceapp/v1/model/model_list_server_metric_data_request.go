package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerMetricDataRequest Request Object
type ListServerMetricDataRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`

	// 服务的命名空间：例如 \"SYS.ECS/AGT.ECS\"，当namespace为AGT.ECS，则查询GPU监控指标：  - SYS.ECS：弹性云服务器的基础监控指标。  - AGT.ECS：弹性云服务器操作系统监控的监控指标（GPU指标）。
	Namespace string `json:"namespace"`

	// 监控查询指标名称:  - SYS.ECS命名空间的指标名称,请参考帮助文档：“[弹性云服务器支持的基础监控指标](https://support.huaweicloud.com/usermanual-ecs/ecs_03_1002.html)”。  - AGT.ECS命名空间的指标名称,请参考帮助文档：“[操作系统监控指标：GPU](https://support.huaweicloud.com/usermanual-ecs/ecs_03_1003.html#section11)”。
	MetricName string `json:"metric_name"`

	// 查询数据起始时间，UNIX时间戳，单位毫秒。
	From string `json:"from"`

	// 查询数据截止时间UNIX时间戳，单位毫秒。from必须小于to。
	To string `json:"to"`

	// 监控数据粒度。 取值范围：  - 1: 实时数据。  - 300: 5分钟粒度。  - 1200: 20分钟粒度。  - 3600: 1小时粒度。  - 14400: 4小时粒度。  - 86400: 1天粒度。
	Period int32 `json:"period"`

	// 数据聚合方式，支持的聚合方式如下:  - average：聚合周期内指标数据的平均值。  - max：聚合周期内指标数据的最大值。  - min：聚合周期内指标数据的最小值。  - sum：聚合周期内指标数据的求和值。  - variance：聚合周期内指标数据的方差。
	Filter string `json:"filter"`
}

func (o ListServerMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ListServerMetricDataRequest", string(data)}, " ")
}
