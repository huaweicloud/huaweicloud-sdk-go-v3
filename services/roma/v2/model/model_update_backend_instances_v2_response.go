package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackendInstancesV2Response Response Object
type UpdateBackendInstancesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的云服务器列表
	Members        *[]VpcMemberInfo `json:"members,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateBackendInstancesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackendInstancesV2Response struct{}"
	}

	return strings.Join([]string{"UpdateBackendInstancesV2Response", string(data)}, " ")
}
