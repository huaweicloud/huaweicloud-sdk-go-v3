package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSpec struct {

	// 规格编码
	SpecCode string `json:"spec_code"`

	// 步数
	StrideNum int32 `json:"stride_num"`
}

func (o CreateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSpec struct{}"
	}

	return strings.Join([]string{"CreateSpec", string(data)}, " ")
}
