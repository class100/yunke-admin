package admin

import (
	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// UpdateConfig 更新配置
func (co *ClientOrg) UpdateConfig(name core.ConfigName, value interface{}, version core.ApiVersion) (config *core.Config, err error) {
	config = new(core.Config)
	err = co.requestApi(core.OrgApiConfigUpdateUrl, gox.HttpMethodPut, value, map[string]string{
		"name": string(name),
	}, version, config)

	return
}
