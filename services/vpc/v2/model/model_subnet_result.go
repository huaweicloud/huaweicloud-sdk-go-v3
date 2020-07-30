/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type SubnetResult struct {
	// uuid形式的一个资源标识。
	Id string `json:"id"`
	// 功能说明：子网的状态。   取值范围：ACTIVE,UNKNOWN,ERROR       ACTIVE表示子网已挂载到ROUTER上       UNKNOWN表示子网还未挂载到ROUTER上       ERROR表示子网状态故障  
	Status string `json:"status"`
}
