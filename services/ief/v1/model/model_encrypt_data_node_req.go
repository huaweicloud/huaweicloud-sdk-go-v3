package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点绑定加密数据配置
type EncryptDataNodeReq struct {

	// 加密数据ID列表
	EncryptDatas []string `json:"encrypt_datas"`
}

func (o EncryptDataNodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataNodeReq struct{}"
	}

	return strings.Join([]string{"EncryptDataNodeReq", string(data)}, " ")
}
