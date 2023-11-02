package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdatableVersionResponse Response Object
type ListUpdatableVersionResponse struct {

	// 集群升级路径总条数
	Count *int32 `json:"count,omitempty"`

	// 集群升级路径列表
	Items          *[]UpdateItemResp `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListUpdatableVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdatableVersionResponse struct{}"
	}

	return strings.Join([]string{"ListUpdatableVersionResponse", string(data)}, " ")
}
