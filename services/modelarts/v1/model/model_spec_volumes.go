package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecVolumes 训练作业挂载卷信息。
type SpecVolumes struct {
	Nfs *Nfs `json:"nfs,omitempty"`

	Pfs *Pfs `json:"pfs,omitempty"`

	Obs *Obs `json:"obs,omitempty"`
}

func (o SpecVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecVolumes struct{}"
	}

	return strings.Join([]string{"SpecVolumes", string(data)}, " ")
}
