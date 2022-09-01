package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApiQuantitiesV2Response struct {

	// API总个数
	InstanceNum *int32 `json:"instance_num,omitempty" xml:"instance_num"`

	// 已发布到release环境的API个数
	NumsOnRelease *int32 `json:"nums_on_release,omitempty" xml:"nums_on_release"`

	// 未发布到release环境的API个数
	NumsOffRelease *int32 `json:"nums_off_release,omitempty" xml:"nums_off_release"`
	HttpStatusCode int    `json:"-"`
}

func (o ListApiQuantitiesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiQuantitiesV2Response struct{}"
	}

	return strings.Join([]string{"ListApiQuantitiesV2Response", string(data)}, " ")
}
