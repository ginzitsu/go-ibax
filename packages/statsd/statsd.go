/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package statsd

import (
	"fmt"
	"strings"

	Client, err = statsd.NewClient(fmt.Sprintf("%s:%d", host, port), name)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	if Client != nil {
		Client.Close()
	}
}

func APIRouteCounterName(method, pattern string) string {
	routeCounterName := strings.Replace(strings.Replace(pattern, ":", "", -1), "/", ".", -1)
	return "api." + strings.ToLower(method) + "." + routeCounterName
}

func DaemonCounterName(daemonName string) string {
	return "daemon." + daemonName
}