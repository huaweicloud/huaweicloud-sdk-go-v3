package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idme/v1/model"
)

type IdmeClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIdmeClient(hcClient *http_client.HcHttpClient) *IdmeClient {
	return &IdmeClient{HcClient: hcClient}
}

func IdmeClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateXdmApplication 创建应用
//
// 本接口用于创建工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) CreateXdmApplication(request *model.CreateXdmApplicationRequest) (*model.CreateXdmApplicationResponse, error) {
	requestDef := GenReqDefForCreateXdmApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateXdmApplicationResponse), nil
	}
}

// CreateXdmApplicationInvoker 创建应用
func (c *IdmeClient) CreateXdmApplicationInvoker(request *model.CreateXdmApplicationRequest) *CreateXdmApplicationInvoker {
	requestDef := GenReqDefForCreateXdmApplication()
	return &CreateXdmApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteXdmApplication 删除应用
//
// 本接口用于删除工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) DeleteXdmApplication(request *model.DeleteXdmApplicationRequest) (*model.DeleteXdmApplicationResponse, error) {
	requestDef := GenReqDefForDeleteXdmApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteXdmApplicationResponse), nil
	}
}

// DeleteXdmApplicationInvoker 删除应用
func (c *IdmeClient) DeleteXdmApplicationInvoker(request *model.DeleteXdmApplicationRequest) *DeleteXdmApplicationInvoker {
	requestDef := GenReqDefForDeleteXdmApplication()
	return &DeleteXdmApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployApplication 部署应用
//
// 本接口用于部署工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) DeployApplication(request *model.DeployApplicationRequest) (*model.DeployApplicationResponse, error) {
	requestDef := GenReqDefForDeployApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployApplicationResponse), nil
	}
}

// DeployApplicationInvoker 部署应用
func (c *IdmeClient) DeployApplicationInvoker(request *model.DeployApplicationRequest) *DeployApplicationInvoker {
	requestDef := GenReqDefForDeployApplication()
	return &DeployApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 获取租户下的应用清单
//
// 本接口用于获取租户在工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）的应用清单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 获取租户下的应用清单
func (c *IdmeClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvs 获取运行服务清单
//
// 本接口用于获取租户在工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）的运行服务清单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) ListEnvs(request *model.ListEnvsRequest) (*model.ListEnvsResponse, error) {
	requestDef := GenReqDefForListEnvs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvsResponse), nil
	}
}

// ListEnvsInvoker 获取运行服务清单
func (c *IdmeClient) ListEnvsInvoker(request *model.ListEnvsRequest) *ListEnvsInvoker {
	requestDef := GenReqDefForListEnvs()
	return &ListEnvsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyApplication 编辑应用
//
// 本接口用于修改工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) ModifyApplication(request *model.ModifyApplicationRequest) (*model.ModifyApplicationResponse, error) {
	requestDef := GenReqDefForModifyApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyApplicationResponse), nil
	}
}

// ModifyApplicationInvoker 编辑应用
func (c *IdmeClient) ModifyApplicationInvoker(request *model.ModifyApplicationRequest) *ModifyApplicationInvoker {
	requestDef := GenReqDefForModifyApplication()
	return &ModifyApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Uninstall 卸载应用
//
// 本接口用于卸载指定运行服务下的工业数字模型驱动引擎（Industrial Digital Model Engine，简称iDME）应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdmeClient) Uninstall(request *model.UninstallRequest) (*model.UninstallResponse, error) {
	requestDef := GenReqDefForUninstall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UninstallResponse), nil
	}
}

// UninstallInvoker 卸载应用
func (c *IdmeClient) UninstallInvoker(request *model.UninstallRequest) *UninstallInvoker {
	requestDef := GenReqDefForUninstall()
	return &UninstallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
