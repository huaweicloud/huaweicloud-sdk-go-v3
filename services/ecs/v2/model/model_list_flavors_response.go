/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ListFlavorsResponse struct {
	// 云服务器规格列表。
	Flavors []Flavor `json:"flavors,omitempty"`
}
