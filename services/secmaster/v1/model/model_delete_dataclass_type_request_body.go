package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataclassTypeRequestBody 删除类型的请求体
type DeleteDataclassTypeRequestBody struct {

	// 是否删除，true删除，false不删除
	IsDelete bool `json:"is_delete"`

	// 删除ids
	Ids []string `json:"ids"`
}

func (o DeleteDataclassTypeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataclassTypeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDataclassTypeRequestBody", string(data)}, " ")
}
