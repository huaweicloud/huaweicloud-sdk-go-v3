package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VisibilityBody 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
type VisibilityBody struct {

	// 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
	Visibility *string `json:"visibility,omitempty"`
}

func (o VisibilityBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VisibilityBody struct{}"
	}

	return strings.Join([]string{"VisibilityBody", string(data)}, " ")
}
