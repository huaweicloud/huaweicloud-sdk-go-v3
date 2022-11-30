package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CompareConfigurationResponse struct {

	// 参数之间的区别集合。
	Differences    *[]DifferentDetails `json:"differences,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CompareConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CompareConfigurationResponse", string(data)}, " ")
}
