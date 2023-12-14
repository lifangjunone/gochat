/**
 * Created by lock
 * Date: 2019-08-12
 * Time: 16:00
 */
package tools

import (
	"fmt"
	"strings"
)

const (
	networkSplit = "@"
)

// ParseNetwork 解析网络地址
func ParseNetwork(str string) (network, addr string, err error) {
	// str = tcp@127.0.0.1:6900
	if idx := strings.Index(str, networkSplit); idx == -1 {
		err = fmt.Errorf("addr: \"%s\" error, must be network@tcp:port or network@unixsocket", str)
		return
	} else {
		// 网络，比如 tcp
		network = str[:idx]
		// 地址： 比如： 127.0.0.1:6900
		addr = str[idx+1:]
		return
	}
}
