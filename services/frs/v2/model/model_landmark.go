package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Landmark struct {
	// 眼睛轮廓，Point为轮廓坐标值。

	EyesContour []Point `json:"eyes_contour"`
	// 嘴巴轮廓，Point为轮廓坐标值。

	MouthContour []Point `json:"mouth_contour"`
	// 人脸轮廓，Point为轮廓坐标值。

	FaceContour []Point `json:"face_contour"`
	// 眉毛轮廓，Point为轮廓坐标值。

	EyebrowContour []map[string]Point `json:"eyebrow_contour"`
	// 鼻子轮廓，Point为轮廓坐标值。

	NoseContour []Point `json:"nose_contour"`
}

func (o Landmark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Landmark struct{}"
	}

	return strings.Join([]string{"Landmark", string(data)}, " ")
}
