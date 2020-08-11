/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 创建弹性公网IP并绑定带宽的请求参数
type CreatePublicipRequestBody struct {
	Bandwidth *CreatePublicipBandwidthOption `json:"bandwidth"`
	// 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。  创建弹性公网IP时，给弹性公网IP绑定企业项目ID。  不指定该参数时，默认值是 0
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	Publicip *CreatePublicipOption `json:"publicip"`
}
