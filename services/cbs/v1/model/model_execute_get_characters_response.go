package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetCharactersResponse Response Object
type ExecuteGetCharactersResponse struct {

	// 配额
	Quota int32 `json:"quota"`

	// 总数
	Total int32 `json:"total"`

	// 偏移
	Offset int32 `json:"offset"`

	// 返回数量
	Count int32 `json:"count"`

	// 形象列表
	Characters     []Character `json:"characters"`
	HttpStatusCode int         `json:"-"`
}

func (o ExecuteGetCharactersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetCharactersResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGetCharactersResponse", string(data)}, " ")
}
