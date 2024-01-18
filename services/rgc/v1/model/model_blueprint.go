package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Blueprint 模板。
type Blueprint struct {

	// 模板ID。
	BlueprintProductId *string `json:"blueprint_product_id,omitempty"`

	// 模板版本。
	BlueprintProductVersion *string `json:"blueprint_product_version,omitempty"`
}

func (o Blueprint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Blueprint struct{}"
	}

	return strings.Join([]string{"Blueprint", string(data)}, " ")
}
