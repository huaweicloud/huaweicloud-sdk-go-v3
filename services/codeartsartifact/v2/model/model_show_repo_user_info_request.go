package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepoUserInfoRequest Request Object
type ShowRepoUserInfoRequest struct {
}

func (o ShowRepoUserInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoUserInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowRepoUserInfoRequest", string(data)}, " ")
}
