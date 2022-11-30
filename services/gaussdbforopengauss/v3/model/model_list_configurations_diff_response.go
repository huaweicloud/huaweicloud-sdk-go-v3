package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConfigurationsDiffResponse struct {

	// 参数组之间的差异集合。
	Differences    *[]ListDiffDetailsResult `json:"differences,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListConfigurationsDiffResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsDiffResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsDiffResponse", string(data)}, " ")
}
