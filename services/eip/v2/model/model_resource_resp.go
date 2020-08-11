/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

type ResourceResp struct {
	// 资源配额对象
	Resources []QuotaShowResp `json:"resources"`
}
