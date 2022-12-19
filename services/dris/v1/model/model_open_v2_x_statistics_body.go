package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：统计数据信息。
type OpenV2XStatisticsBody struct {
	Source *StatisticsSourceDto `json:"source,omitempty"`

	// **参数说明**：数据上报的时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'。 例如：2021-01-08T02:03:41Z。
	Time *string `json:"time,omitempty"`

	// **参数说明**：统计周期，单位秒。
	Period *int32 `json:"period,omitempty"`

	// **参数说明**：道路路的角度，区分道路方向，向东为0度，逆时针增加。
	Direction float32 `json:"direction,omitempty"`

	// **参数说明**：统计周期内的车辆数。
	Flow *int32 `json:"flow,omitempty"`

	// **参数说明**：车辆平均速度，单位km/h。
	AverageSpeed float32 `json:"average_speed,omitempty"`

	// **参数说明**：设备编码。 **取值范围**：长度不超过64，只允许字母、数字、以及_等字符的组合。
	Esn *string `json:"esn,omitempty"`

	// **参数说明**：路口id，对应到一组雷视拟合设备，检测一个特定的路口或者路段。
	CrossId *string `json:"cross_id,omitempty"`

	// **参数说明**：路段的交通流方向，交通流方向按照“西北规则”进行定义，即尽量选择西北的点作为正向起点，先西后北。西北规则具体说明请参见 [“西北规则”说明](此处添加support文档的url)。 0：正向 1：逆向 2：正向转逆向的连接线 3：逆向转正向的连接线 9：为方向未确定
	TrafficDirection *int32 `json:"traffic_direction,omitempty"`

	// **参数说明**：道路特征，0为主路，1为汇入匝道，2为汇出匝道，3为辅道
	RoadKind *int32 `json:"road_kind,omitempty"`

	// **参数说明**：不同车辆类型的流量统计。
	VehicleClassFlow *[]ModelFlow `json:"vehicle_class_flow,omitempty"`

	// **参数说明**：分车道统计的占有率列表。
	Occupancy *[]LaneOccupancy `json:"occupancy,omitempty"`
}

func (o OpenV2XStatisticsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenV2XStatisticsBody struct{}"
	}

	return strings.Join([]string{"OpenV2XStatisticsBody", string(data)}, " ")
}
