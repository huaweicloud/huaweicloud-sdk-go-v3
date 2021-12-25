package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRepositoryStatusResponse struct {
	Error *Error `json:"error,omitempty"`
	// 响应结果

	Result *int32 `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryStatusResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryStatusResponse", string(data)}, " ")
}
