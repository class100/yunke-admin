package admin

import (
	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// UpdateConfig 更新配置
func (oc *OrgClient) UpdateConfig(name yunke.ConfigName, value interface{}, version yunke.ApiVersion) (config *yunke.Config, err error) {
	config = new(yunke.Config)
	err = oc.requestApi(yunke.OrgApiConfigUpdateUrl, gox.HttpMethodPut, value, map[string]string{
		"name": string(name),
	}, version, config)

	return
}
