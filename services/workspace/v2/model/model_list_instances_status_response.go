package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesStatusResponse Response Object
type ListInstancesStatusResponse struct {

	// 统计信息。
	Statics        *[]InstanceStatusStatistics `json:"statics,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListInstancesStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesStatusResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesStatusResponse", string(data)}, " ")
}
