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

func checkAliyunToken() bool {
	tokenId := viper.GetString(global.ConfAliyunTokenIdKey)
	tokenKey := viper.GetString(global.ConfAliyunTokenSecretKey)

	return len(tokenId) >= 1 && len(tokenKey) >= 1
}
func checkDomainConfig() bool {
	domain := viper.GetString(global.ConfDomainKey)
	subDomain := viper.GetString(global.ConfSubDomainKey)

	return validator.IsValidDomain(domain) && len(subDomain) >= 1
}
