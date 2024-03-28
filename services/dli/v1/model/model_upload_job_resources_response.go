package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadJobResourcesResponse Response Object
type UploadJobResourcesResponse struct {

	// 模块名。
	GroupName *string `json:"group_name,omitempty"`

	// 上传分组资源状态。
	Status *string `json:"status,omitempty"`

	// 该模块包含的资源包名列表。
	Resources *[]string `json:"resources,omitempty"`

	// 模块上传的unix时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 模块更新的unix时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 本次上传是同步还是异步
	IsAsync *bool `json:"is_async,omitempty"`

	// 资源包拥有者
	Owner *string `json:"owner,omitempty"`

	// 分组资源包的详细信息。
	Details        *[]UploadJobResourcesDetail `json:"details,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UploadJobResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadJobResourcesResponse struct{}"
	}

	return strings.Join([]string{"UploadJobResourcesResponse", string(data)}, " ")
}
