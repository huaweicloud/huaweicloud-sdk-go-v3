package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationDiffReqV3 struct {
	DiffPara *ParaGroupDiff `json:"diff_para,omitempty"`
}

func (o ConfigurationDiffReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDiffReqV3 struct{}"
	}

	return strings.Join([]string{"ConfigurationDiffReqV3", string(data)}, " ")
}
