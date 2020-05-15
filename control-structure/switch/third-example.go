package main

/*
	define a generic type - i interface{}
*/
func checkType(i interface{}) string {
	switch i.(type) {
	case int:
		return "type: integer"
	case float32, float64:
		return "type: real"
	case string:
		return "type: string"
	case func():
		return "type: function"
	default:
		return "type: unknow"
	}
}
