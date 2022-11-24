package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectIterationsV4Response struct {

	// 迭代总数
	Total *int32 `json:"total,omitempty"`

	// 迭代信息
	Iterations     *[]ListProjectVersionsV4ResponseBodyIterations `json:"iterations,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o ListProjectIterationsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectIterationsV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectIterationsV4Response", string(data)}, " ")
}
