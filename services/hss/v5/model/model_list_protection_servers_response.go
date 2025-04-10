package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectionServersResponse Response Object
type ListProtectionServersResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// list
	DataList       *[]ProtectionServeInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListProtectionServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionServersResponse struct{}"
	}

	return strings.Join([]string{"ListProtectionServersResponse", string(data)}, " ")
}
