package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type OsfederationProtocol struct {

	// 协议ID。
	Id string `json:"id" xml:"id"`
}

func (o OsfederationProtocol) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsfederationProtocol struct{}"
	}

	return strings.Join([]string{"OsfederationProtocol", string(data)}, " ")
}
