package model

import (
	"encoding/json"

	"strings"
)

// 镜像元素
type ImageData struct {
	//
	ImageId *string `json:"image_id,omitempty"`
}

func (o ImageData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageData struct{}"
	}

	return strings.Join([]string{"ImageData", string(data)}, " ")
}
