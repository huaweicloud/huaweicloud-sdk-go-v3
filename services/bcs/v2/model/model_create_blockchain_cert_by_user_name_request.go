package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBlockchainCertByUserNameRequest Request Object
type CreateBlockchainCertByUserNameRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id"`

	// peer组织名称
	OrgName string `json:"org_name"`

	// 用户名称，字符串长度4-24，仅支持小写字母和数字，以小写字母开头
	UserName string `json:"user_name"`

	Body *CreateBlockchainCertByUserNameRequestBody `json:"body,omitempty"`
}

func (o CreateBlockchainCertByUserNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBlockchainCertByUserNameRequest struct{}"
	}

	return strings.Join([]string{"CreateBlockchainCertByUserNameRequest", string(data)}, " ")
}
