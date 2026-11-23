// Package mydispatcher Package dispatcher implement the rate limiter and the online device counter
package mydispatcher

import "github.com/xtls/xray-core/features/routing"

//go:generate go run github.com/xtls/xray-core/common/errors/errorgen

// Type returns routing.DispatcherType() so mydispatcher serves as the active core dispatcher.
func Type() interface{} {
	return routing.DispatcherType()
}
