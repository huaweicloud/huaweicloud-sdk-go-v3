package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAnonymousAuthRandomResponse struct {

	// 下一跳URL。
	SiteUrl *string `json:"siteUrl,omitempty" xml:"siteUrl"`

	// 鉴权随机数。
	Random         *string `json:"random,omitempty" xml:"random"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAnonymousAuthRandomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnonymousAuthRandomResponse struct{}"
	}

	return strings.Join([]string{"CreateAnonymousAuthRandomResponse", string(data)}, " ")
}
