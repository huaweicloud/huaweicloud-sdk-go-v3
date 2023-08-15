package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareConfigurationResponse Response Object
type CompareConfigurationResponse struct {

	// 参数模板之间的区别集合。
	Differences    *[]DiffDetails `json:"differences,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CompareConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CompareConfigurationResponse", string(data)}, " ")
}
