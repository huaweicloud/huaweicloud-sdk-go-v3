package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryJobInstancesByNameResponse Response Object
type ListFactoryJobInstancesByNameResponse struct {

	// 总记录数
	Total *int64 `json:"total,omitempty"`

	// 作业实例状态
	Instances      *[]JobInstance `json:"instances,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListFactoryJobInstancesByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryJobInstancesByNameResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryJobInstancesByNameResponse", string(data)}, " ")
}
