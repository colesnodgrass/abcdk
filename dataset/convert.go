package dataset

// ConvertIntToFloat64 recursively converts integer-like values to float64 in a map[string]any
func ConvertIntToFloat64(input map[string]any) map[string]any {
	output := make(map[string]any)

	for key, value := range input {
		switch v := value.(type) {
		case int:
			output[key] = float64(v)
		case int8:
			output[key] = float64(v)
		case int16:
			output[key] = float64(v)
		case int32:
			output[key] = float64(v)
		case int64:
			output[key] = float64(v)
		case uint:
			output[key] = float64(v)
		case uint8:
			output[key] = float64(v)
		case uint16:
			output[key] = float64(v)
		case uint32:
			output[key] = float64(v)
		case uint64:
			output[key] = float64(v)
		case map[string]any:
			output[key] = ConvertIntToFloat64(v)
		case []any:
			convertedSlice := make([]any, len(v))
			for i, item := range v {
				if m, ok := item.(map[string]any); ok {
					convertedSlice[i] = ConvertIntToFloat64(m)
				} else {
					switch itemValue := item.(type) {
					case int:
						convertedSlice[i] = float64(itemValue)
					case int8:
						convertedSlice[i] = float64(itemValue)
					case int16:
						convertedSlice[i] = float64(itemValue)
					case int32:
						convertedSlice[i] = float64(itemValue)
					case int64:
						convertedSlice[i] = float64(itemValue)
					case uint:
						convertedSlice[i] = float64(itemValue)
					case uint8:
						convertedSlice[i] = float64(itemValue)
					case uint16:
						convertedSlice[i] = float64(itemValue)
					case uint32:
						convertedSlice[i] = float64(itemValue)
					case uint64:
						convertedSlice[i] = float64(itemValue)
					default:
						convertedSlice[i] = item
					}
				}
			}
			output[key] = convertedSlice
		default:
			output[key] = value
		}
	}

	return output
}
