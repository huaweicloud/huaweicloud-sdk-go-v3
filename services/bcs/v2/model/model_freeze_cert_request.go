package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type FreezeCertRequest struct {

	// userName
	UserName string `json:"user_name"`

	// blockchainID
	BlockchainId string `json:"blockchain_id"`

	// orgName
	OrgName string `json:"org_name"`
}

func (o FreezeCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeCertRequest struct{}"
	}

	return strings.Join([]string{"FreezeCertRequest", string(data)}, " ")
}
