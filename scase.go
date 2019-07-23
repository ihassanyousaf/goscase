package goscase

import (
	"encoding/json"
	"strings"
	"unicode"
)

func snakeToCamel(value string) string {
	if value[0] == '_' {
		return value
	}
	var result []rune
	flag := false
	value = strings.TrimSpace(value)
	if value[len(value)-1] == '_' {
		value = value[:len(value)-1]
	}
	for _, perItem := range value {
		if perItem == 95 {
			flag = true
		} else if flag {
			flag = false
			result = append(result, unicode.ToUpper(perItem))
		} else {
			result = append(result,perItem)
		}
	}
	return string(result)
}

func camelToSnake(value string) string {
	if value[0] == '_' {
		return value
	} else if unicode.IsUpper(rune(value[0]))   {
		return value
	}
	var result []rune
	value = strings.TrimSpace(value)
	for _, perItem := range value {
		if unicode.IsUpper(perItem) {
			result = append(result, 95)
			result = append(result, unicode.ToLower(perItem))
		} else {
			result = append(result,perItem)
		}
	}
	return string(result)
}


func converted(key string, camel bool) (string, bool) {
	result := ""
	if camel {
		result = snakeToCamel(key)
	} else {
		result= camelToSnake(key)
	}
	if result == key {
		return result, false
	}
	return result, true
}

func arrayParsing(input []interface{}, flag bool)  []interface{} {
	for  _ , val := range input {
		switch val.(type) {
		case map[string]interface{}:
			mapParsing(val.(map[string]interface{}), flag)
		case []interface{}:
			arrayParsing(val.([]interface{}), flag)
		default:
		}
	}
	return input
}

func mapParsing(input map[string]interface{}, flag bool) map[string]interface{} {
	for key, val := range input {
		if newKey, changed := converted(key,flag); changed{
			input[newKey] = input[key]
			delete(input, key)
		}
		switch val.(type) {
		case map[string]interface{}:
				mapParsing(val.(map[string]interface{}), flag)
		case []interface{}:
				arrayParsing(val.([]interface{}), flag)
		default:
		}
	}
	return input
}


func snakeToCamelJSON(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	byteData, _ := json.Marshal(data);
	_ = json.Unmarshal(byteData, &result)
	return mapParsing(result, true)
}

func camelToSnakeJSON(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	byteData, _ := json.Marshal(data);
	_ = json.Unmarshal(byteData, &result)
	return mapParsing(result, false)
}
