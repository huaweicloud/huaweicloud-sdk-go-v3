package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VehicleCertificateResult struct {

	// 合格证编号。
	CertificateNumber *string `json:"certificate_number,omitempty"`

	// 发证日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 车辆制造企业名称。
	ManufactureName *string `json:"manufacture_name,omitempty"`

	// 车辆品牌。
	VehicleBrand *string `json:"vehicle_brand,omitempty"`

	// 车辆名称。
	VehicleName *string `json:"vehicle_name,omitempty"`

	// 车辆型号。
	VehicleModel *string `json:"vehicle_model,omitempty"`

	// 车架号。
	Vin *string `json:"vin,omitempty"`

	// 车身颜色。
	VehicleColor *string `json:"vehicle_color,omitempty"`

	// 底盘型号。
	ChassisModel *string `json:"chassis_model,omitempty"`

	// 底盘ID。
	ChassisId *string `json:"chassis_id,omitempty"`

	// 底盘合格证编号。
	ChassisCertificateNumber *string `json:"chassis_certificate_number,omitempty"`

	// 发动机型号。
	EngineModel *string `json:"engine_model,omitempty"`

	// 发动机号。
	EngineNumber *string `json:"engine_number,omitempty"`

	// 燃料种类。
	FuelType *string `json:"fuel_type,omitempty"`

	// 排量。
	Displacement *string `json:"displacement,omitempty"`

	// 功率。
	Power *string `json:"power,omitempty"`

	// 排放标准。
	EmissionStandard *string `json:"emission_standard,omitempty"`

	// 油耗。
	FuelConsumption *string `json:"fuel_consumption,omitempty"`

	// 外廓尺寸-长。
	OverallDimensionLength *string `json:"overall_dimension_length,omitempty"`

	// 外廓尺寸-宽。
	OverallDimensionWidth *string `json:"overall_dimension_width,omitempty"`

	// 外廓尺寸-高。
	OverallDimensionHeight *string `json:"overall_dimension_height,omitempty"`

	// 货厢内部尺寸-长。
	ContainerDimensionLength *string `json:"container_dimension_length,omitempty"`

	// 货厢内部尺寸-宽。
	ContainerDimensionWidth *string `json:"container_dimension_width,omitempty"`

	// 货厢内部尺寸-高。
	ContainerDimensionHeight *string `json:"container_dimension_height,omitempty"`

	// 钢板弹簧片数。
	SpringQuantity *string `json:"spring_quantity,omitempty"`

	// 轮胎数。
	TireQuantity *string `json:"tire_quantity,omitempty"`

	// 轮胎规格。
	TireSize *string `json:"tire_size,omitempty"`

	// 轮距-前。
	FrontWheelTrack *string `json:"front_wheel_track,omitempty"`

	// 轮距-后。
	RearWheelTrack *string `json:"rear_wheel_track,omitempty"`

	// 轴距。
	Wheelbase *string `json:"wheelbase,omitempty"`

	// 轴荷。
	AxleLoad *string `json:"axle_load,omitempty"`

	// 轴数。
	AxleQuantity *string `json:"axle_quantity,omitempty"`

	// 转向形式。
	SteeringForm *string `json:"steering_form,omitempty"`

	// 总质量。
	TotalWeight *string `json:"total_weight,omitempty"`

	// 整备质量。
	EquipmentWeight *string `json:"equipment_weight,omitempty"`

	// 额定载质量。
	MaximumLadenMass *string `json:"maximum_laden_mass,omitempty"`

	// 载质量利用系数。
	MassUtilizationCoefficient *string `json:"mass_utilization_coefficient,omitempty"`

	// 准牵引总质量。
	TractionWeight *string `json:"traction_weight,omitempty"`

	// 半挂车鞍座最大允许总质量。
	MaximumLoadMass *string `json:"maximum_load_mass,omitempty"`

	// 驾驶室准乘人数。
	CabPassengerCapacity *string `json:"cab_passenger_capacity,omitempty"`

	// 额定载客。
	PassengerCapacity *string `json:"passenger_capacity,omitempty"`

	// 最高设计车速。
	MaxDesignSpeed *string `json:"max_design_speed,omitempty"`

	// 车辆制造日期。
	ManufactureDate *string `json:"manufacture_date,omitempty"`

	// 字段的置信度，取值范围0~1。 置信度越大，本次识别的字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o VehicleCertificateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleCertificateResult struct{}"
	}

	return strings.Join([]string{"VehicleCertificateResult", string(data)}, " ")
}
