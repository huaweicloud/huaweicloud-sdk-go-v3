package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObsBucketAclResponse Response Object
type ShowObsBucketAclResponse struct {

	// 桶风险
	BucketRisk *string `json:"bucket_risk,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	// 桶ACL
	ObsAcl         *string `json:"obs_acl,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowObsBucketAclResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObsBucketAclResponse struct{}"
	}

	return strings.Join([]string{"ShowObsBucketAclResponse", string(data)}, " ")
}
