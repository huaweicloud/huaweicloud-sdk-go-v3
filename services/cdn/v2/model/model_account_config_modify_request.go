package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountConfigModifyRequest 租户配置权限修改请求体。
type AccountConfigModifyRequest struct {

	// OBS委托授权，使用OBS私有桶作为源站时需要开启该委托授权。当前仅支持传on：开启OBS委托授权。
	ObsAgencyStatus *string `json:"obs_agency_status,omitempty"`

	// SCM委托授权，配置HTTPS证书时，证书来源选择SCM证书时需要开启该委托授权。当前仅支持传on：开启SCM委托授权。
	ScmAgencyStatus *string `json:"scm_agency_status,omitempty"`
}

func (o AccountConfigModifyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountConfigModifyRequest struct{}"
	}

	return strings.Join([]string{"AccountConfigModifyRequest", string(data)}, " ")
}
