package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryJobDependInstancesResponse Response Object
type ShowFactoryJobDependInstancesResponse struct {

	// 查询作业上下游依赖关系响应体。
	Body           *[]ShowFactoryJobDependInstancesResponseBody `json:"body,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ShowFactoryJobDependInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryJobDependInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowFactoryJobDependInstancesResponse", string(data)}, " ")
}
