package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowColumnInfoResultResponse Response Object
type ShowColumnInfoResultResponse struct {

	// 指定数据库表列信息
	Results *[]DbObjectColumnInfo `json:"results,omitempty"`

	// 列表中的项目总数，与分页无关
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowColumnInfoResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowColumnInfoResultResponse struct{}"
	}

	return strings.Join([]string{"ShowColumnInfoResultResponse", string(data)}, " ")
}
