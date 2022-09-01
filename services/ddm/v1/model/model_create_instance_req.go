package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateInstanceReq struct {
	Instance *CreateInstanceDetail `json:"instance" xml:"instance"`

	ExtendParam *CreateInstanceExtendParam `json:"extend_param,omitempty" xml:"extend_param"`
}

func (o CreateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceReq", string(data)}, " ")
}
