package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateInternetAccessOption 公网访问
type TemplateInternetAccessOption struct {
	Publicip *TemplatePublicIpOption `json:"publicip,omitempty"`
}

func (o TemplateInternetAccessOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateInternetAccessOption struct{}"
	}

	return strings.Join([]string{"TemplateInternetAccessOption", string(data)}, " ")
}
