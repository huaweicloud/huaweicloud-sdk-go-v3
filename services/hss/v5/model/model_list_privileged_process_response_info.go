package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPrivilegedProcessResponseInfo struct {

	// data list
	DataList *[]PrivilegedProcessResponseInfo `json:"data_list,omitempty"`

	// total number
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o ListPrivilegedProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivilegedProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"ListPrivilegedProcessResponseInfo", string(data)}, " ")
}
