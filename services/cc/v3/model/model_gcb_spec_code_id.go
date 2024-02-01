package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GcbSpecCodeId 功能说明：线路规格编码UUID。
type GcbSpecCodeId struct {

	// 功能说明：线路规格编码UUID。
	SpecCodeId *string `json:"spec_code_id,omitempty"`
}

func (o GcbSpecCodeId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbSpecCodeId struct{}"
	}

	return strings.Join([]string{"GcbSpecCodeId", string(data)}, " ")
}
