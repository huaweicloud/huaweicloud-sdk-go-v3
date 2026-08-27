package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyObjectDetailInfo 应用对象名称信息
type ApplyObjectDetailInfo struct {

	// id
	Id *string `json:"id,omitempty"`

	// 对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 对象名称
	ObjectName *string `json:"object_name,omitempty"`

	// 域名称（用户）
	ObjectDomain *string `json:"object_domain,omitempty"`
}

func (o ApplyObjectDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyObjectDetailInfo struct{}"
	}

	return strings.Join([]string{"ApplyObjectDetailInfo", string(data)}, " ")
}
