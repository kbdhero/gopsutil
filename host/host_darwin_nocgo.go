// +build darwin
// +build !cgo

package host

import "github.com/kbdhero/gopsutil/internal/common"

func SensorsTemperatures() ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
