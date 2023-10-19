package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsDifferencesResponse Response Object
type ListConfigurationsDifferencesResponse struct {

	// 参数之间的区别集合。
	Differences    *[]ParamGroupParameterDifferences `json:"differences,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListConfigurationsDifferencesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsDifferencesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsDifferencesResponse", string(data)}, " ")
}
