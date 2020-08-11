/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type ChangeSeversOsMetadata struct {
	// 重装云服务器过程中注入用户数据。  支持注入文本、文本文件或gzip文件。注入内容最大长度32KB。注入内容，需要进行base64格式编码。
	UserData string `json:"user_data,omitempty"`
}
