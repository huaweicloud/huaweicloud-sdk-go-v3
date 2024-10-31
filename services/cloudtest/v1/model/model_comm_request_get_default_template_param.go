package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestGetDefaultTemplateParam struct {
	Params *GetDefaultTemplateParam `json:"params,omitempty"`
}

func (o CommRequestGetDefaultTemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestGetDefaultTemplateParam struct{}"
	}

	return strings.Join([]string{"CommRequestGetDefaultTemplateParam", string(data)}, " ")
}
