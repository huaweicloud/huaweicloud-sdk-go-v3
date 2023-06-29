package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMeasureUnitsRequest Request Object
type ListMeasureUnitsRequest struct {

	// 语言。zh_CN：中文en_US：英文缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListMeasureUnitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeasureUnitsRequest struct{}"
	}

	return strings.Join([]string{"ListMeasureUnitsRequest", string(data)}, " ")
}
