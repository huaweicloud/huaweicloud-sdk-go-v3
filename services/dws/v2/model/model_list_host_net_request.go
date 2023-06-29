package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostNetRequest Request Object
type ListHostNetRequest struct {

	// 集群ID。获取方法，请参见9.6-获取集群ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 数据条目数。
	Limit int32 `json:"limit"`

	// 数据偏移量。
	Offset int32 `json:"offset"`
}

func (o ListHostNetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostNetRequest struct{}"
	}

	return strings.Join([]string{"ListHostNetRequest", string(data)}, " ")
}
