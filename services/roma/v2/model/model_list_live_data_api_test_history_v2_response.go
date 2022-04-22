package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLiveDataApiTestHistoryV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的测试结果对象列表
	Histories      *[]LdApiTestHistoryInfoV2 `json:"histories,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListLiveDataApiTestHistoryV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataApiTestHistoryV2Response struct{}"
	}

	return strings.Join([]string{"ListLiveDataApiTestHistoryV2Response", string(data)}, " ")
}
