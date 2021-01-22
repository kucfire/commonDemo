package log

import "time"

func getYear() int {

}

func Init() {
	pathVariableTable = make(map[byte]func(*time.Time) int, 5)
	pathVariableTable['Y'] = getYear

}
