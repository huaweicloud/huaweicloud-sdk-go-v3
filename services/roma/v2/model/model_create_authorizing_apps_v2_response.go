package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAuthorizingAppsV2Response struct {
	// API与APP的授权关系列表

	Auths          *[]ApiAuthRelations `json:"auths,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateAuthorizingAppsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizingAppsV2Response struct{}"
	}

	return strings.Join([]string{"CreateAuthorizingAppsV2Response", string(data)}, " ")
}
