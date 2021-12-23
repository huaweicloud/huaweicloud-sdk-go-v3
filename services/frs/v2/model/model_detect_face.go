package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DetectFace struct {
	BoundingBox *BoundingBox `json:"bounding_box"`

	Attributes *Attributes `json:"attributes,omitempty"`

	Landmark *Landmark `json:"landmark,omitempty"`
}

func (o DetectFace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFace struct{}"
	}

	return strings.Join([]string{"DetectFace", string(data)}, " ")
}
