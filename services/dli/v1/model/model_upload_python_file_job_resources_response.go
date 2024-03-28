package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPythonFileJobResourcesResponse Response Object
type UploadPythonFileJobResourcesResponse struct {

	// \"UPLOADING\"表示正在上传。 \"READY\"表示模块包已上传。 \"FAILED\"表示模块包上传失败。
	Status *string `json:"status,omitempty"`

	// 资源模块描述。
	Description *string `json:"description,omitempty"`

	// 该资源模块包含的资源包名列表。
	Resources *[]string `json:"resources,omitempty"`

	// 资源模块上传的unix时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 资源模块更新的unix时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 模块名。
	GroupName *string `json:"group_name,omitempty"`

	// 资源包拥有者
	Owner *string `json:"owner,omitempty"`

	// 是否使用异步方式上传资源包。默认值为“false”，表示不使用异步方式。推荐使用异步方式上传资源包。
	IsAsync *bool `json:"is_async,omitempty"`

	// 分组资源包的详细信息
	Details *[]UploadJobResourcesDetail `json:"details,omitempty"`

	// 资源模块名
	ModuleName *string `json:"module_name,omitempty"`

	// 资源模块类型。jar:用户jar文件;pyFile:用户python文件;file:用户文件
	ModuleType     *string `json:"module_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadPythonFileJobResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPythonFileJobResourcesResponse struct{}"
	}

	return strings.Join([]string{"UploadPythonFileJobResourcesResponse", string(data)}, " ")
}
