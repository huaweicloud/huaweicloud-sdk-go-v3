package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectivityTaskRequestBody 测试地址连通性 {   \"address\": \"iam.cn-north-7.ulanqab.huawei.com:443\" }
type CreateConnectivityTaskRequestBody struct {

	// 测试地址
	Address string `json:"address"`
}

func (o CreateConnectivityTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectivityTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConnectivityTaskRequestBody", string(data)}, " ")
}
