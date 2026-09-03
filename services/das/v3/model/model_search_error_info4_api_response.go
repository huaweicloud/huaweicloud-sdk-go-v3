package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchErrorInfo4ApiResponse Response Object
type SearchErrorInfo4ApiResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// binlog解析错误信息列表
	ErrorTransInfos *[]ErrorTransInfo `json:"error_trans_infos,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o SearchErrorInfo4ApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchErrorInfo4ApiResponse struct{}"
	}

	return strings.Join([]string{"SearchErrorInfo4ApiResponse", string(data)}, " ")
}
