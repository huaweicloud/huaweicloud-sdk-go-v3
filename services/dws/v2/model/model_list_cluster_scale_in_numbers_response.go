package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterScaleInNumbersResponse Response Object
type ListClusterScaleInNumbersResponse struct {

	// 参数解释: 合适的缩容数。 取值范围: 不涉及。
	ShrinkSequence *[]string `json:"shrink_sequence,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListClusterScaleInNumbersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterScaleInNumbersResponse struct{}"
	}

	return strings.Join([]string{"ListClusterScaleInNumbersResponse", string(data)}, " ")
}
