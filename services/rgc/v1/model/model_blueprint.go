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

	// 模板部署参数。
	Variables *string `json:"variables,omitempty"`

	// 模板是否包含多账号资源。
	IsBlueprintHasMultiAccountResource *bool `json:"is_blueprint_has_multi_account_resource,omitempty"`
}

func (o Blueprint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Blueprint struct{}"
	}

	return strings.Join([]string{"Blueprint", string(data)}, " ")
}
