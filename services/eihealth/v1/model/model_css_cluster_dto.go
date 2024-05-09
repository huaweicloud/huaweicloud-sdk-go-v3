package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CssClusterDto css集群详情
type CssClusterDto struct {

	// 已绑定的集群id
	Id *string `json:"id,omitempty"`

	// css集群名称
	Name *string `json:"name,omitempty"`

	// css集群总存储
	Storage *int32 `json:"storage,omitempty"`

	// css集群导入时间
	ImportTime *string `json:"import_time,omitempty"`
}

func (o CssClusterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CssClusterDto struct{}"
	}

	return strings.Join([]string{"CssClusterDto", string(data)}, " ")
}
