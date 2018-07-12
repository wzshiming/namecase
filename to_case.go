package namecase

import (
	"unsafe"
)

const (
	pace      = 'a' - 'A'
	underline = '_'
	strike    = '-'
	space     = ' '
)

// ToUpperSpace to XX YY ZZ
func ToUpperSpace(s string) string {
	return toSplit(s, space, true)
}

// ToLowerSpace to xx yy zz
func ToLowerSpace(s string) string {
	return toSplit(s, space, false)
}

// ToUpperStrike to XX-YY-ZZ
func ToUpperStrike(s string) string {
	return toSplit(s, strike, true)
}

// ToLowerStrike to xx-yy-zz
func ToLowerStrike(s string) string {
	return toSplit(s, strike, false)
}

// ToUpperSnake to XX_YY_ZZ
func ToUpperSnake(s string) string {
	return toSplit(s, underline, true)
}

// ToLowerSnake to xx_yy_zz
func ToLowerSnake(s string) string {
	return toSplit(s, underline, false)
}

// ToPascal to XxYyZz
func ToPascal(s string) string {
	return toHump(s, true)
}

// ToUpperHump to XxYyZz
func ToUpperHump(s string) string {
	return toHump(s, true)
}

// ToCamel to xxYyZz
func ToCamel(s string) string {
	return toHump(s, false)
}

// ToLowerHump to xxYyZz
func ToLowerHump(s string) string {
	return toHump(s, false)
}

// toSplit to xx yy zz
func toSplit(s string, sep byte, upper bool) string {
	data := make([]byte, 0, len(s)*2)
	pu := false // 上一个是下划线
	pb := false // 上一个是大写
	cu := true  // 当前是下划线
	cb := false // 当前是大写
	cs := false // 当前是小写
	iu := false // 前面忽略了一个下划线
	ib := false // 以进入正式的内容
	for i := 0; i != len(s); i++ {
		v := s[i]

		pu = cu
		pb = cb

		cs = v >= 'a' && v <= 'z'
		cb = !cs && v >= 'A' && v <= 'Z'
		cu = !(cb || cs)

		// 去除无用的前缀
		if !ib {
			if cu {
				if v == underline || v == strike || v == space {
					continue
				}
			}
			ib = true
		}

		// 当前是下划线
		if cu {
			if v == underline || v == strike || v == space {
				iu = true // 忽略下划线
				continue
			}
			data = append(data, v)
			continue
		}

		// 大小写 转换
		if upper {
			if cs {
				v -= pace
			}
		} else {
			if cb {
				v += pace
			}
		}

		if (cb && !pb && !pu) || //  当前是大写 上一个不是大写 上一个不是下划线
			(iu && !cu) { // 之前忽略了一个下划线 当前不是下划线
			iu = false
			data = append(data, sep)
		}

		data = append(data, v)
	}

	return *(*string)(unsafe.Pointer(&data))
}

// toHump to xxYyZz
func toHump(s string, upper bool) string {
	data := make([]byte, 0, len(s))
	iu := true // 之前有下划线
	// ps := false // 上一个是小写
	pb := false // 上一个是大写
	cs := false // 当前是小写
	cb := false // 当前是大写
	for i := 0; i != len(s); i++ {
		v := s[i]

		// ps = cs
		pb = cb
		cs = v >= 'a' && v <= 'z'
		cb = !cs && v >= 'A' && v <= 'Z'

		if !(cs || cb) {
			if v == underline || v == strike || v == space {
				iu = true
				continue
			}
			data = append(data, v)
			continue
		}

		if iu {
			if cs { // 之前有下划线 当前是小写
				v -= pace
			}
			iu = false
		} else if pb && cb { // 上一个是大写 当前是大写
			v += pace
		}
		data = append(data, v)
	}

	if !upper && len(data) != 0 {
		data[0] += pace
	}
	return *(*string)(unsafe.Pointer(&data))
}
