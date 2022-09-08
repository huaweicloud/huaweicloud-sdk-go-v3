package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UnfreezeCertRequest struct {

	// userName
	UserName string `json:"user_name"`

	// blockchainID
	BlockchainId string `json:"blockchain_id"`

	// orgName
	OrgName string `json:"org_name"`

	Body *UnfreezeCertRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UnfreezeCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeCertRequest struct{}"
	}

	return strings.Join([]string{"UnfreezeCertRequest", string(data)}, " ")
}
