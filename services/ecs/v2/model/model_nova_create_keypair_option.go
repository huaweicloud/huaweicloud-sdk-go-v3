/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

//  
type NovaCreateKeypairOption struct {
	// 导入的公钥信息。  建议导入的公钥长度不大于1024字节。  说明：  - 长度超过1024字节会导致云服务器注入该密钥失败。
	PublicKey string `json:"public_key,omitempty"`
	// 密钥类型，值为“ssh”或“x509”。  说明：  - 微版本2.2支持。
	Type string `json:"type,omitempty"`
	// 密钥名称。  新创建的密钥名称不能和已有密钥名称相同。
	Name string `json:"name"`
	// 密钥的用户ID。  说明：  - 微版本2.10支持。
	UserId string `json:"user_id,omitempty"`
}
