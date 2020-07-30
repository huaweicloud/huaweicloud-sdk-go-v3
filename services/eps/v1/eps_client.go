package v1

import (
    http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1/model"
)

type EpsClient struct {
    hcClient *http_client.HcHttpClient
}

func NewEpsClient(hcClient *http_client.HcHttpClient) *EpsClient {
    return &EpsClient{hcClient: hcClient}
}

func EpsClientBuilder() *http_client.HcHttpClientBuilder {
    builder := http_client.NewHcHttpClientBuilder()
    return builder
}


//创建企业项目。
func (c *EpsClient) CreateEP(request *model.CreateEPRequest) (*model.CreateEPResponse, error) {
    requestDef := GenReqDefForCreateEP(request)
    resp, responseDef := GenRespForCreateEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//停用企业项目。
func (c *EpsClient) DisableEP(request *model.DisableEPRequest) (*model.DisableEPResponse, error) {
    requestDef := GenReqDefForDisableEP(request)
    resp, responseDef := GenRespForDisableEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//启用企业项目。
func (c *EpsClient) EnableEP(request *model.EnableEPRequest) (*model.EnableEPResponse, error) {
    requestDef := GenReqDefForEnableEP(request)
    resp, responseDef := GenRespForEnableEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询企业项目的API版本列表。
func (c *EpsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
    requestDef := GenReqDefForListApiVersions(request)
    resp, responseDef := GenRespForListApiVersions()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询当前用户已授权的企业项目列表，用户可以使用企业项目绑定资源。
func (c *EpsClient) ListEP(request *model.ListEPRequest) (*model.ListEPResponse, error) {
    requestDef := GenReqDefForListEP(request)
    resp, responseDef := GenRespForListEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//迁移资源到目标企业项目。
func (c *EpsClient) MigrateResource(request *model.MigrateResourceRequest) (*model.MigrateResourceResponse, error) {
    requestDef := GenReqDefForMigrateResource(request)
    resp, responseDef := GenRespForMigrateResource()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//修改企业项目。当前仅支持修改名称和描述。
func (c *EpsClient) ModifyEP(request *model.ModifyEPRequest) (*model.ModifyEPResponse, error) {
    requestDef := GenReqDefForModifyEP(request)
    resp, responseDef := GenRespForModifyEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询指定的企业项目API版本号详情
func (c *EpsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
    requestDef := GenReqDefForShowApiVersion(request)
    resp, responseDef := GenRespForShowApiVersion()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询企业项目详情。
func (c *EpsClient) ShowEP(request *model.ShowEPRequest) (*model.ShowEPResponse, error) {
    requestDef := GenReqDefForShowEP(request)
    resp, responseDef := GenRespForShowEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询企业项目的配额信息。
func (c *EpsClient) ShowEPQuota(request *model.ShowEPQuotaRequest) (*model.ShowEPQuotaResponse, error) {
    requestDef := GenReqDefForShowEPQuota(request)
    resp, responseDef := GenRespForShowEPQuota()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询企业项目下绑定的资源详情。
func (c *EpsClient) ShowResourceBindEP(request *model.ShowResourceBindEPRequest) (*model.ShowResourceBindEPResponse, error) {
    requestDef := GenReqDefForShowResourceBindEP(request)
    resp, responseDef := GenRespForShowResourceBindEP()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

