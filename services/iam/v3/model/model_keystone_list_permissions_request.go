package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListPermissionsRequest struct {
	// 权限名，获取方式请参见：[获取权限名、权限ID](https://support.huaweicloud.com/api-iam/iam_10_0001.html)。

	Name *string `json:"name,omitempty"`
	// 账号ID，获取方式请参见：[获取账号ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。    > - 如果填写此参数，则返回账号下所有自定义策略。 > - 如果不填写此参数，则返回所有系统权限（包含系统策略和系统角色）。

	DomainId *string `json:"domain_id,omitempty"`
	// 分页查询时数据的页数，查询值最小为1，默认值为1。需要与per_page同时使用。

	Page *int32 `json:"page,omitempty"`
	// 分页查询时每页的数据个数，取值范围为[1,300], 默认值为300。需要与page同时使用。

	PerPage *int32 `json:"per_page,omitempty"`
}

func (o KeystoneListPermissionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListPermissionsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListPermissionsRequest", string(data)}, " ")
}
