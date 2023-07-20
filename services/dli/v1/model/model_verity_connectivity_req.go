package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerityConnectivityReq 测试地址连通性 {   \"address\": \"iam.cn-north-7.ulanqab.huawei.com:443\" }
type VerityConnectivityReq struct {

	// 测试地址
	Address string `json:"address"`
}

func (o VerityConnectivityReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerityConnectivityReq struct{}"
	}

	return strings.Join([]string{"VerityConnectivityReq", string(data)}, " ")
}
