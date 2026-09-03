package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlLimitUserInstanceRequest Request Object
type ListSqlLimitUserInstanceRequest struct {
	Body *ListSqlLimitUserInstanceRequestBody `json:"body,omitempty"`
}

func (o ListSqlLimitUserInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitUserInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListSqlLimitUserInstanceRequest", string(data)}, " ")
}
