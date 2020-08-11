/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ListQuotasResponse struct {
	Quotas *ResourceResp `json:"quotas,omitempty"`
}
