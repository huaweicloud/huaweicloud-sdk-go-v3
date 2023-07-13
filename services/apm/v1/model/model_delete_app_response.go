package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppResponse Response Object
type DeleteAppResponse struct {

	// 返回ok的结果，表示删除成功。
	Ok             *string `json:"ok,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppResponse", string(data)}, " ")
}
