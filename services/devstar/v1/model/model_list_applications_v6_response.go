package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicationsV6Response struct {

	// 应用列表
	Applications *[]ApplicationV3 `json:"applications,omitempty" xml:"applications"`

	// 应用列表总条数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListApplicationsV6Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsV6Response struct{}"
	}

	return strings.Join([]string{"ListApplicationsV6Response", string(data)}, " ")
}
