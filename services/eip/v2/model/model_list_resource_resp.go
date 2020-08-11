/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 资源实例详情
type ListResourceResp struct {
	// 资源详情。 资源对象，用于扩展。默认为空
	ResouceDetail map[string]interface{} `json:"resouce_detail,omitempty"`
	// 资源ID
	ResourceId string `json:"resource_id,omitempty"`
	// 资源名称，没有默认为空字符串
	ResourceName string `json:"resource_name,omitempty"`
	// 标签列表，没有标签默认为空数组
	Tags []ResourceTagResp `json:"tags,omitempty"`
}
