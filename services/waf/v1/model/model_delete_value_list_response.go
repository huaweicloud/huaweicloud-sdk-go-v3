package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteValueListResponse struct {
	// 引用表id

	Id *string `json:"id,omitempty"`
	// 引用表名称

	Name *string `json:"name,omitempty"`
	// 引用表类型

	Type *string `json:"type,omitempty"`
	// 引用表描述

	Description *string `json:"description,omitempty"`
	// 引用表时间戳

	Timestamp *int64 `json:"timestamp,omitempty"`
	// 引用表的值

	Values         *[]string `json:"values,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteValueListResponse struct{}"
	}

	return strings.Join([]string{"DeleteValueListResponse", string(data)}, " ")
}
