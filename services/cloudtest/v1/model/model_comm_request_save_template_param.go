package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestSaveTemplateParam struct {
	Params *SaveTemplateParam `json:"params,omitempty"`
}

func (o CommRequestSaveTemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestSaveTemplateParam struct{}"
	}

	return strings.Join([]string{"CommRequestSaveTemplateParam", string(data)}, " ")
}
