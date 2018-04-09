package namecase

import (
	"unsafe"
)

const (
	pace = 'a' - 'A'
	wall = '_'
)

func ToUpperSnake(s string) string {
	return toSnake(s, true)
}

func ToLowerSnake(s string) string {
	return toSnake(s, false)
}

// ToSnake
func toSnake(s string, upper bool) string {
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
		cu = v == wall

		// 去除无用的前缀
		if !ib {
			if cu {
				continue
			}
			ib = true
		}

		// 当前是下划线
		if cu {
			iu = true // 忽略下划线
			continue
		}

		cb = v >= 'A' && v <= 'Z'

		// 大小写 转换
		if upper {
			cs = !cb && v >= 'a' && v <= 'z'
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
			data = append(data, wall)
		}

		data = append(data, v)
	}

	return *(*string)(unsafe.Pointer(&data))
}

func ToUpperHump(s string) string {
	return toHump(s, true)
}

func ToLowerHump(s string) string {
	return toHump(s, false)
}

// ToHump
func toHump(s string, upper bool) string {
	data := make([]byte, 0, len(s))
	iu := true // 之前有下划线
	// ps := false // 上一个是小写
	pb := false // 上一个是大写
	cs := false // 当前是小写
	cb := false // 当前是大写
	for i := 0; i != len(s); i++ {
		v := s[i]
		if v == wall {
			iu = true
			continue
		}

		// ps = cs
		pb = cb
		cs = v >= 'a' && v <= 'z'
		cb = !cs && v >= 'A' && v <= 'Z'

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
