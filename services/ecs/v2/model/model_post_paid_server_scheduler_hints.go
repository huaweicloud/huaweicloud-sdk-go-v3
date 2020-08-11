/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type PostPaidServerSchedulerHints struct {
	// 云服务器组ID，UUID格式。
	Group string `json:"group,omitempty"`
	// 专属主机的ID。专属主机的ID仅在tenancy为dedicated时生效。
	DedicatedHostId string `json:"dedicated_host_id,omitempty"`
	// 在指定的专属主机或者共享主机上创建弹性云服务器云主机。参数值为shared或者dedicated。
	Tenancy string `json:"tenancy,omitempty"`
}
