package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachImageServerAppResponse Response Object
type AttachImageServerAppResponse struct {

	// 分发软件信息的URI。
	Uri            *string `json:"uri,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachImageServerAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachImageServerAppResponse struct{}"
	}

	return strings.Join([]string{"AttachImageServerAppResponse", string(data)}, " ")
}
