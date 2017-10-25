package ligo

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
