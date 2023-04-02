package Calculator

import (
	Operations "SupUmbrela/Utils"
	Logger "SupUmbrela/Utils/Log"
	"fmt"
	"strconv"
)

func ToFloat32(value *string) float32 {
	v, err := strconv.ParseFloat(*value, 32)
	if err != nil {
		Logger.TraceLog("Failed to parse value", Logger.ERRO, fmt.Sprintf("Value: %s", *value))
	}
	return float32(v)
}

var mapOp map[uint32]func(*MathString) float32 = map[uint32]func(*MathString) float32{
	Operations.Hash("+"): add,
}

type MathString struct {
	lOp     interface{}
	rOp     interface{}
	operand string
}

func NewOperation(left interface{}, operand string, right interface{}) MathString {
	var mString MathString = MathString{
		lOp: left, rOp: right, operand: operand,
	}
	return mString
}

func (ms *MathString) Calculate() float32 {
	return 0
}

func (ms *MathString) getFloats() (float32, float32) {
	lVal, ok := ms.lOp.(float32)
	if !ok {
		Logger.TraceLog("Failed to cast to float32 on first value", Logger.ERRO)
	}
	rVal, ok := ms.rOp.(float32)
	if !ok {
		Logger.TraceLog("Failed to cast to float32 on second value", Logger.ERRO)
	}

	return lVal, rVal
}

func add(ms *MathString) float32 {
	lVal, rVal := ms.getFloats()
	return lVal + rVal
}
