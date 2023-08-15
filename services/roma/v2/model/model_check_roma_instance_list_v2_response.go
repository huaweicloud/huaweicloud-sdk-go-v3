package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRomaInstanceListV2Response Response Object
type CheckRomaInstanceListV2Response struct {

	// 列表总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 实例列表
	Instances      *[]RomaInstanceCheckListRespInstances `json:"instances,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o CheckRomaInstanceListV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRomaInstanceListV2Response struct{}"
	}

	return strings.Join([]string{"CheckRomaInstanceListV2Response", string(data)}, " ")
}
