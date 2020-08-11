/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// This is a auto create Body Object
type NovaCreateServersRequestBody struct {
	Server *NovaCreateServersOption `json:"server"`
	OsschedulerHints *NovaCreateServersSchedulerHint `json:"os:scheduler_hints,omitempty"`
}
