package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckStarrocksParamsResponse Response Object
type CheckStarrocksParamsResponse struct {

	// 参数之间区别的集合。
	Differences    *[]ParamGroupParameterDifferences `json:"differences,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o CheckStarrocksParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckStarrocksParamsResponse struct{}"
	}

	return strings.Join([]string{"CheckStarrocksParamsResponse", string(data)}, " ")
}
