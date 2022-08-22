/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package initializator

import (
	"github.com/AghostPrj/ddns/internal/global"
	"github.com/dchest/validator"
	"github.com/spf13/viper"
)

func checkDomainServiceProviderConf() bool {
	provider := viper.GetString(global.ConfDomainServiceProvider)

	if len(provider) < 1 {
		return false
	}

	switch provider {
	case global.DomainServiceProviderAliyun:
		return checkAliyunToken()
	case global.DomainServiceProviderDnspod:
		return checkDnsPodToken()
	default:
		return false
	}
}

func checkDnsPodToken() bool {
	tokenSecret := viper.GetString(global.ConfDnspodTokenSecretKey)

	return len(tokenSecret) >= 1
}

func checkAliyunToken() bool {
	tokenId := viper.GetString(global.ConfAliyunTokenIdKey)
	tokenSecret := viper.GetString(global.ConfAliyunTokenSecretKey)

	return len(tokenId) >= 1 && len(tokenSecret) >= 1
}
func checkDomainConfig() bool {
	domain := viper.GetString(global.ConfDomainKey)
	subDomain := viper.GetString(global.ConfSubDomainKey)

	return validator.IsValidDomain(domain) && len(subDomain) >= 1
}
func checkIpSource() bool {
	publicIpSource := viper.GetString(global.ConfPublicIpSourceKey)

	return publicIpSource == global.PublicIpSourceInterface ||
		publicIpSource == global.PublicIpSourceIpLookupService

}
