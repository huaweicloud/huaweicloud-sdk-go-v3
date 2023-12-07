package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Blueprint 蓝图。
type Blueprint struct {

	// 蓝图ID。
	BlueprintProductId *string `json:"blueprint_product_id,omitempty"`

	// 蓝图版本。
	BlueprintProductVersion *string `json:"blueprint_product_version,omitempty"`
}

func (o Blueprint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Blueprint struct{}"
	}

	return strings.Join([]string{"Blueprint", string(data)}, " ")
}
