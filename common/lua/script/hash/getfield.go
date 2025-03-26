package luaHash

import (
	"os"
	"runtime"
)

type getField struct {
	function string
}

func (h *getField) Name() string     { return "hash_getfield" }
func (h *getField) Function() string { return h.function }

var getFieldScript *getField

func init() {
	_, path, _, _ := runtime.Caller(0)
	res, err := os.ReadFile(path + "getfield.lua")
	if err != nil {
		panic(err.Error())
	}

	getFieldScript = &getField{string(res)}
}

func GetGetField() *getField {
	return getFieldScript
}

var HGetFieldNotExists = "FieldNotExists"
var HGetTableNotExists = "TableNotExists"
