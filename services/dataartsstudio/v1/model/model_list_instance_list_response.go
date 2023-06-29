package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceListResponse Response Object
type ListInstanceListResponse struct {

	// api操作对应的实例列表
	Instances      *[]InstanceForApiActionDto `json:"instances,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListInstanceListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceListResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceListResponse", string(data)}, " ")
}
