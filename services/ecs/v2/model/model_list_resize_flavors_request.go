/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ListResizeFlavorsRequest struct {
	InstanceUuid string `json:"instance_uuid,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	SortDir string `json:"sort_dir,omitempty"`
	SortKey string `json:"sort_key,omitempty"`
	SourceFlavorId string `json:"source_flavor_id,omitempty"`
	SourceFlavorName string `json:"source_flavor_name,omitempty"`
}
