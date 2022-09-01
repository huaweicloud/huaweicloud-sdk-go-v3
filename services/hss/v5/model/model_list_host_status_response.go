package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostStatusResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty" xml:"total_num"`

	// 查询弹性云服务器状态列表
	DataList       *[]Host `json:"data_list,omitempty" xml:"data_list"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHostStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostStatusResponse struct{}"
	}

	return strings.Join([]string{"ListHostStatusResponse", string(data)}, " ")
}
