package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPropertyValuesResponse Response Object
type ShowPropertyValuesResponse struct {

	// 查询设备的属性值
	Properties     *[]PropertyValue `json:"properties,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowPropertyValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPropertyValuesResponse struct{}"
	}

	return strings.Join([]string{"ShowPropertyValuesResponse", string(data)}, " ")
}
