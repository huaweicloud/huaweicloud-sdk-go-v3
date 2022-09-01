package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Attributes struct {
	Dress *Dress `json:"dress" xml:"dress"`

	// 是否带眼镜： • yes：带眼镜 • dark：带墨镜 • none：未戴眼镜 • unknown：未知
	Glass string `json:"glass" xml:"glass"`

	// 是否戴帽子： • yes：戴帽子 • none：未戴帽子 • unknown：未知
	Hat string `json:"hat" xml:"hat"`

	// 年龄。
	Age int32 `json:"age" xml:"age"`

	// 是否戴口罩： • yes：戴口罩 • none：未戴口罩 • unknown：未知
	Mask string `json:"mask" xml:"mask"`

	// 胡须： • yes：有胡须 • none：无胡须 • unknown：未知
	Beard string `json:"beard" xml:"beard"`

	// 图片类型： • idcard：证件照 • monitor：摄像头监控 • internet photo：网络图片
	Phototype string `json:"phototype" xml:"phototype"`

	Quality *FaceQuality `json:"quality" xml:"quality"`

	// 发型： • long：长发 • short：短发 • unknown：未知
	Hair string `json:"hair" xml:"hair"`

	Expression *AttributesExpression `json:"expression" xml:"expression"`

	// 人脸图片旋转角（顺时针偏转角度），支持0°、90°、180°和270°图片旋转。
	FaceAngle int32 `json:"face_angle" xml:"face_angle"`
}

func (o Attributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attributes struct{}"
	}

	return strings.Join([]string{"Attributes", string(data)}, " ")
}
