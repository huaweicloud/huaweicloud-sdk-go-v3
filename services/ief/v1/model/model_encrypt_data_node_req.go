package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncryptDataNodeReq 边缘节点绑定加密数据配置
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
