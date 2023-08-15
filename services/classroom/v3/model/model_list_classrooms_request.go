package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClassroomsRequest Request Object
type ListClassroomsRequest struct {

	// 信息记录的起始编号
	Offset *int32 `json:"offset,omitempty"`

	// 每页包含的信息记录数
	Limit *int32 `json:"limit,omitempty"`

	// 查询的课堂类别，默认查询所有的课堂。 取值范围： create：只查询当前用户创建的课堂。 attend：只查询当前用户加入的课堂。 all：查询当前用户所有的课堂。
	QueryType *string `json:"query_type,omitempty"`
}

func (o ListClassroomsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClassroomsRequest struct{}"
	}

	return strings.Join([]string{"ListClassroomsRequest", string(data)}, " ")
}
