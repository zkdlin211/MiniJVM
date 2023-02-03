package test

import (
	"MiniJVM/src/runtime"
	"reflect"
	"testing"
)

func TestOperandStack(t *testing.T) {
	frame := runtime.NewFrame(100, 100)
	testLocalVars(frame.LocalVars(), t)
	testOperandStack(frame.OperandStack(), t)
}

func testLocalVars(vars runtime.LocalVars, t *testing.T) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)

	assertEqual(t, vars.GetInt(0), 100)
	assertEqual(t, vars.GetInt(1), -100)
	assertEqual(t, vars.GetLong(2), 2997924580)
	assertEqual(t, vars.GetLong(4), -2997924580)
	assertEqual(t, vars.GetFloat(6), 3.1415926)
	assertEqual(t, vars.GetDouble(7), 2.71828182845)
	assertEqual(t, vars.GetRef(9), nil)
}

func testOperandStack(stack *runtime.OperandStack, t *testing.T) {
	stack.PushInt(100)
	stack.PushInt(-100)
	stack.PushLong(2997924580)
	stack.PushLong(-2997924580)
	stack.PushFloat(3.1415926)
	stack.PushDouble(2.71828182845)
	stack.PushRef(nil)
	assertEqual(t, stack.PopRef(), nil)
	assertEqual(t, stack.PopDouble(), 2.71828182845)
	assertEqual(t, stack.PopFloat(), 3.1415926)
	assertEqual(t, stack.PopLong(), -2997924580)
	assertEqual(t, stack.PopLong(), 2997924580)
	assertEqual(t, stack.PopInt(), -100)
	assertEqual(t, stack.PopInt(), 100)
}
func assertEqual(t *testing.T, actual, expected interface{}) {
	if expected == nil {
		if actual != nil {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
		return
	}
	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)

	if !actualValue.Type().ConvertibleTo(expectedValue.Type()) {
		t.Errorf("Expected value %v to be convertible to type %v, but it wasn't", actual, expectedValue.Type())
	} else {
		if actualValue.Convert(expectedValue.Type()).Interface() != expected {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	}
}
