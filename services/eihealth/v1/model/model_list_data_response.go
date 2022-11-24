package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDataResponse struct {

	// 数据对象（目录，文件）总数量
	Count *int64 `json:"count,omitempty"`

	// 数据对象列表
	Datas *[]DataRsp `json:"datas,omitempty"`

	// 下一页开始标签
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataResponse struct{}"
	}

	return strings.Join([]string{"ListDataResponse", string(data)}, " ")
}
