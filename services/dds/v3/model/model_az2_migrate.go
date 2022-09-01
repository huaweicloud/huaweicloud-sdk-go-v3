package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规格信息。
type Az2Migrate struct {

	// 可用区ID。
	Code string `json:"code" xml:"code"`

	// 可用区描述。
	Description string `json:"description" xml:"description"`

	// 当前可用区的的状态。 - ENABLED，表示该可用区（组合）可用。 - DISABLED，表示该可用区（组合）不可用。
	Status string `json:"status" xml:"status"`
}

func (o Az2Migrate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Az2Migrate struct{}"
	}

	return strings.Join([]string{"Az2Migrate", string(data)}, " ")
}
