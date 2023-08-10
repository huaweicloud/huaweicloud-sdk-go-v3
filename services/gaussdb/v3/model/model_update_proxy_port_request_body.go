package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyPortRequestBody 修改读写分离端口号请求体。
type UpdateProxyPortRequestBody struct {

	// 修改后的读写分离端口。  GaussDB(for MySQL) Proxy端口号范围：大于等于1025，小于等于65534，不包含端口1033、5342-5345、12017、20000、20201、20202、33062、33071。
	Port int32 `json:"port"`
}

func (o UpdateProxyPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyPortRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProxyPortRequestBody", string(data)}, " ")
}
