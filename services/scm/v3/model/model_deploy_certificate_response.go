package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployCertificateResponse Response Object
type DeployCertificateResponse struct {

	// 部署失败的资源列表。
	FailureList    *[]FaiureResource `json:"failure_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeployCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeployCertificateResponse", string(data)}, " ")
}
