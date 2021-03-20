package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddApplicationResponse struct {
	// 资源空间ID，唯一标识一个资源空间，由物联网平台在创建资源空间时分配。资源空间对应的是物联网平台原有的应用，在物联网平台的含义与应用一致，只是变更了名称。

	AppId *string `json:"app_id,omitempty"`
	// 资源空间名称。

	AppName *string `json:"app_name,omitempty"`
	// 资源空间创建时间，格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。

	CreateTime *string `json:"create_time,omitempty"`
	// 是否为默认资源空间

	DefaultApp *bool `json:"default_app,omitempty"`
	// app的类型，标准版：Junior | 高级版：Normal

	AppType *string `json:"app_type,omitempty"`
	// 用户名。

	Username *string `json:"username,omitempty"`
	// app与用户的授权关系时，响应为：all | bind | edit | query ，其中bind权限类似于ALL权限，属于子用户权限。

	Permission *string `json:"permission,omitempty"`
	// 迁移前实例ID。

	LastInstanceId *string `json:"last_instance_id,omitempty"`
	// 当前实例ID。

	CurrentInstanceId *string `json:"current_instance_id,omitempty"`
	// 对接的服务名

	ServiceName *string `json:"service_name,omitempty"`
	// 是否冻结

	Freezed        *bool `json:"freezed,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o AddApplicationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddApplicationResponse struct{}"
	}

	return strings.Join([]string{"AddApplicationResponse", string(data)}, " ")
}
