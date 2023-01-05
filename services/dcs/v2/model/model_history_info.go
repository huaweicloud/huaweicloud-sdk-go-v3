package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryInfo struct {

	// 修改记录ID
	HistoryId *string `json:"history_id,omitempty"`

	// 修改类型
	Type *string `json:"type,omitempty"`

	// 修改时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改状态
	Status *string `json:"status,omitempty"`
}

func (o HistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryInfo struct{}"
	}

	return strings.Join([]string{"HistoryInfo", string(data)}, " ")
}
