package ligo

import "fmt"

func KeySearch(key string, data interface{}, allowInterfaces bool) interface{} {
	keyedData := data.(map[string]interface{})
	var val interface{}
	for k, v := range keyedData {
		if allowInterfaces {
			if key == k {
				val = v
				break
			}
		} else {
			switch vv := v.(type) {
			case map[string]interface{}:
				val = KeySearch(key, v, allowInterfaces)
				if val != nil {
					break
				}
			default:
				if key == k {
					val = vv
					break
				}
			}
		}
		if val != nil {
			break
		}
	}
	return val
}

func CheckType(inter interface{}) {
	switch v := inter.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case float64:
		fmt.Printf("Float64: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case map[string]interface{}:
		fmt.Printf("Interface: %v\n", v)
	default:
		// And here I'm feeling dumb. ;)
		fmt.Printf("Unknown type")
	}
}
