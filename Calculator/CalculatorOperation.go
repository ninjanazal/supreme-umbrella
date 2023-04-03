package Calculator

import (
	Operations "SupUmbrela/Utils"
	Logger "SupUmbrela/Utils/Log"
)

// mapOp is a map that stores the supported mathematical operations as hash codes and their corresponding functions.
// Each function takes a pointer to a MathString struct as input and returns the result
// of the mathematical operation as a float32 value.
var mapOp map[uint32]func(*MathString) float32 = map[uint32]func(*MathString) float32{
	Operations.Hash("+"): add,
	Operations.Hash("-"): subtract,
	Operations.Hash("*"): multiply,
	Operations.Hash("/"): divide,
}

// MathString is a struct that holds the information of a mathematical expression.
// lOp (interface{}): The left operand of the expression.
// rOp (interface{}): The right operand of the expression.
// operand (string): The operand that links both operands.
type MathString struct {
	lOp     interface{}
	rOp     interface{}
	operand string
}

// NewOperation creates a new MathString object with the given left and right operands and the specified operand.
// left (interface{}): The left operand of the operation.
// operand (string): The operand to perform between the left and right operands.
// right (interface{}): The right operand of the operation.
// Return {MathString}: A new MathString object with the specified operands and operand.
func NewOperation(left interface{}, operand string, right interface{}) MathString {
	var mString MathString = MathString{
		lOp: left, rOp: right, operand: operand,
	}
	return mString
}

// Calculate evaluates the mathematical expression represented by the MathString object.
// Return {float32}: The result of the mathematical expression.
// If there is an error during the evaluation, 0.0 is returned, and an error message is logged.
func (ms *MathString) Calculate() float32 {
	switch ms.lOp.(type) {
	case float32:
		switch ms.rOp.(type) {
		case MathString:
			mExp, ok := ms.rOp.(MathString)
			if ok {
				ms.rOp = mExp.Calculate()
			}
			Logger.TraceLog("Failed to process left operation ", Logger.ERRO)

		}
	case MathString:
		mExp, ok := ms.lOp.(MathString)
		if ok {
			ms.lOp = mExp.Calculate()
		}
		Logger.TraceLog("Failed to process right operation ", Logger.ERRO)
	}
	opFunc, ok := mapOp[Operations.Hash(ms.operand)]
	if ok {
		return opFunc(ms)
	}

	Logger.TraceLog("Failed to process operator "+string(ms.operand), Logger.ERRO)
	return 0.0
}

// getFloats returns the left and right operands of a MathString as float32 values.
// Return (float32, float32): A tuple of two float32 values
// representing the left and right operands of the MathString.
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

// add adds the left and right values in the MathString and returns the result.
// ms (*MathString): A pointer to the MathString to perform the operation on.
// Return {float32}: The sum of the left and right values.
func add(ms *MathString) float32 {
	lVal, rVal := ms.getFloats()
	return lVal + rVal
}

// subtract subtracts the right value from the left value in the MathString and returns the result.
// ms (*MathString): A pointer to the MathString to perform the operation on.
// Return {float32}: The difference between the left and right values.
func subtract(ms *MathString) float32 {
	lVal, rVal := ms.getFloats()
	return lVal - rVal
}

// multiply multiplies the left and right values of the given MathString.
// ms (*MathString): A pointer to the MathString to operate on.
// Return {float32}: The result of multiplying the left and right values.
func multiply(ms *MathString) float32 {
	lVal, rVal := ms.getFloats()
	return lVal * rVal
}

// divide divides the left and right operands of the given MathString.
// It returns the result of the division as a float32.
// ms (*MathString): A pointer to the MathString object containing the operands to divide.
// Return {float32}: The result of the division.
func divide(ms *MathString) float32 {
	lVal, rVal := ms.getFloats()
	return lVal / rVal
}
