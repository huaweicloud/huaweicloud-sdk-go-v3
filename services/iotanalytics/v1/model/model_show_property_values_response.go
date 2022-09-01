package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPropertyValuesResponse struct {

	// 查询设备的属性值
	Properties     *[]PropertyValue `json:"properties,omitempty" xml:"properties"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowPropertyValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPropertyValuesResponse struct{}"
	}

	return strings.Join([]string{"ShowPropertyValuesResponse", string(data)}, " ")
}
