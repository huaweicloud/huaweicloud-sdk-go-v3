/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type NovaListKeypairsResponse struct {
	// 密钥信息列表。
	Keypairs []NovaListKeypairsResult `json:"keypairs,omitempty"`
}
