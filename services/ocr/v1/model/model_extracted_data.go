package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ExtractedData struct {
	MathInfo *MathInfo `json:"math_info,omitempty"`
}

func (o ExtractedData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtractedData struct{}"
	}

	return strings.Join([]string{"ExtractedData", string(data)}, " ")
}
