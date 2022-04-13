package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateBlockchainCertByUserNameRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`
	// peer组织名称

	OrgName string `json:"org_name"`
	// 用户名称，字符串长度4-24，必须包含a-z，0-9，以小写字母开头，以小写字母或者数字结尾

	UserName string `json:"user_name"`
}

func (o CreateBlockchainCertByUserNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBlockchainCertByUserNameRequest struct{}"
	}

	return strings.Join([]string{"CreateBlockchainCertByUserNameRequest", string(data)}, " ")
}
