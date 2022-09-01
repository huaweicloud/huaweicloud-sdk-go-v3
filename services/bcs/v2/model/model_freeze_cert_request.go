package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type FreezeCertRequest struct {

	// userName
	UserName string `json:"user_name" xml:"user_name"`

	// blockchainID
	BlockchainId string `json:"blockchain_id" xml:"blockchain_id"`

	// orgName
	OrgName string `json:"org_name" xml:"org_name"`

	Body *FreezeCertRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o FreezeCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeCertRequest struct{}"
	}

	return strings.Join([]string{"FreezeCertRequest", string(data)}, " ")
}
