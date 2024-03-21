package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdmParamVoVersionUpdateViewDtoMultiViewModelViewUpdateAttrDto struct {
	Params *VersionUpdateViewDtoMultiViewModelViewUpdateAttrDto `json:"params,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"applicationId,omitempty"`
}

func (o RdmParamVoVersionUpdateViewDtoMultiViewModelViewUpdateAttrDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdmParamVoVersionUpdateViewDtoMultiViewModelViewUpdateAttrDto struct{}"
	}

	return strings.Join([]string{"RdmParamVoVersionUpdateViewDtoMultiViewModelViewUpdateAttrDto", string(data)}, " ")
}
