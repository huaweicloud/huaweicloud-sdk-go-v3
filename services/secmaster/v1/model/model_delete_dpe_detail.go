package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDpeDetail 删除分类映射响应详细信息
type DeleteDpeDetail struct {

	// 删除成功id集合
	SuccessIdList *[]string `json:"success_id_list,omitempty"`

	// 删除失败id集合
	ErrorIdList *[]string `json:"error_id_list,omitempty"`
}

func (o DeleteDpeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDpeDetail struct{}"
	}

	return strings.Join([]string{"DeleteDpeDetail", string(data)}, " ")
}
