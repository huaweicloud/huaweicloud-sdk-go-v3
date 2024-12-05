package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateBindingGeipRequestBody struct {

	// 带宽包id
	GcbId *string `json:"gcb_id,omitempty"`

	GlobalEips *[]BindingGeipBody `json:"global_eips,omitempty"`
}

func (o CreateBindingGeipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindingGeipRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBindingGeipRequestBody", string(data)}, " ")
}
