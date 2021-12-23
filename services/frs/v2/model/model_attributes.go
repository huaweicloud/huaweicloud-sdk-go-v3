package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Attributes struct {
	Dress *Dress `json:"dress"`
	// 是否带眼镜： • yes：带眼镜 • dark：带墨镜 • none：未戴眼镜 • unknown：未知

	Glass string `json:"glass"`
	// 性别： • male：男 • female：女 • unknown：未知

	Gender string `json:"gender"`
	// 围绕Y轴旋转，偏航角，范围[-180,180]。

	YawAngle float64 `json:"yaw_angle"`
	// 围绕Z轴旋转，翻滚角，范围[-180,180]。

	RollAngle float64 `json:"roll_angle"`
	// 围绕X轴旋转，俯仰角，范围[-180,180]。

	PitchAngle float64 `json:"pitch_angle"`
	// 是否戴帽子： • yes：戴帽子 • none：未戴帽子 • unknown：未知

	Hat string `json:"hat"`
	// 人脸轮廓坐标值。

	Headpose []float64 `json:"headpose"`
	// 年龄。

	Age int32 `json:"age"`
	// 笑脸。

	Smile string `json:"smile"`
	// 是否戴口罩： • yes：戴口罩 • none：未戴口罩 • unknown：未知

	Mask string `json:"mask"`
	// 胡须： • yes：有胡须 • none：无胡须 • unknown：未知

	Beard string `json:"beard"`
	// 肤色： • brown：棕 • yellow：黄 • white：白 • black：黑 • unknown：未知

	Skin string `json:"skin"`
	// 民族： • han：汉族 • other：其他 • unknown：未知

	Ethnic string `json:"ethnic"`
	// 图片类型： • idcard：证件照 • monitor：摄像头监控 • internet photo：网络图片

	Phototype string `json:"phototype"`

	Quality *FaceQuality `json:"quality"`
	// 发型： • long：长发 • short：短发 • unknown：未知

	Hair string `json:"hair"`

	Expression *AttributesExpression `json:"expression"`
	// 人脸图片旋转角（顺时针偏转角度），支持0°、90°、180°和270°图片旋转。

	FaceAngle int32 `json:"face_angle"`
}

func (o Attributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attributes struct{}"
	}

	return strings.Join([]string{"Attributes", string(data)}, " ")
}
