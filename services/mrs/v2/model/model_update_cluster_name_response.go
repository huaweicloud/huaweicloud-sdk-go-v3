package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterNameResponse Response Object
type UpdateClusterNameResponse struct {

	// 更新映射请求操作结果，succeeded为操作成功，failed为操作失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClusterNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterNameResponse", string(data)}, " ")
}
