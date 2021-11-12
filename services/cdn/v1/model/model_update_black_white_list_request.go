package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateBlackWhiteListRequest struct {
	// 需要设置IP黑白名单的域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
	// 当用户开启企业项目功能时，该参数生效，表示资源所属项目，不传表示查询默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *BlackWhiteListBody `json:"body,omitempty"`
}

func (o UpdateBlackWhiteListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListRequest", string(data)}, " ")
}
