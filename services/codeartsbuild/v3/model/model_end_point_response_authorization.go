package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndPointResponseAuthorization 授权
type EndPointResponseAuthorization struct {

	// 参数
	Parameters *interface{} `json:"parameters,omitempty"`

	// 计划
	Scheme *interface{} `json:"scheme,omitempty"`
}

func (o EndPointResponseAuthorization) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndPointResponseAuthorization struct{}"
	}

	return strings.Join([]string{"EndPointResponseAuthorization", string(data)}, " ")
}
