/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ListResizeFlavorsResponse struct {
	// 云服务器规格列表。
	Flavors []ListResizeFlavorsResult `json:"flavors,omitempty"`
}
