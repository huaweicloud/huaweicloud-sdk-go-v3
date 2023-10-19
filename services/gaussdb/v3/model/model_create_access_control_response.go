package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessControlResponse Response Object
type CreateAccessControlResponse struct {

	// 是否已开启访问控制。 取值： - true：开启。 - false：关闭。
	OpenAccessControl *bool `json:"open_access_control,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o CreateAccessControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessControlResponse struct{}"
	}

	return strings.Join([]string{"CreateAccessControlResponse", string(data)}, " ")
}
