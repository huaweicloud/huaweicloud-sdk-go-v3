package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessKeyReqBody The request body of create access key request.
type CreateAccessKeyReqBody struct {

	// 访问密钥描述信息。
	Description *string `json:"description,omitempty"`
}

func (o CreateAccessKeyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessKeyReqBody struct{}"
	}

	return strings.Join([]string{"CreateAccessKeyReqBody", string(data)}, " ")
}
