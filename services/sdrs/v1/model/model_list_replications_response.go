package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListReplicationsResponse struct {

	// 复制对列表。
	Replications *[]ShowReplicationParams `json:"replications,omitempty" xml:"replications"`

	// 列表中包含的复制对个数。
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListReplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListReplicationsResponse", string(data)}, " ")
}
