package main

import (
	"reflect"
	"strings"
)

// Register(
//    Function{name:"funA", argumentTypes:[]string{"Boolean", "Integer"}, isVariadic:false},
//    Function{name:"funB", argumentTypes:[]string{"Integer"}, isVariadic:false},
//    Function{name:"funC", argumentTypes:[]string{"Integer"}, isVariadic:true}
// )
// FindMatches("Boolean", "Integer")            -> [funA]
// FindMatches("Integer")                       -> [funB, funC]
// FindMatches("Integer", "Integer", "Integer") -> [funC]

type libraryFormat string

const (
	listFormat libraryFormat = "list"
	mapFormat  libraryFormat = "map"
)

type FunctionService interface {
	Register(functions ...Function)
	FindMatches(argumentTypes ...string) []string
}

type Function struct {
	name          string
	argumentTypes []string
	isVariadic    bool
}

type FunctionListLibrary struct {
	funcsList []Function
}

func NewFunctionLibrary(mode libraryFormat) FunctionService {
	switch mode {
	case listFormat:
		return &FunctionListLibrary{}
	case mapFormat:
		return &FunctionMapLibrary{}
	}
	return nil
}

func (fl *FunctionListLibrary) Register(functions ...Function) {
	fl.funcsList = append(fl.funcsList, functions...)
}

func (fl *FunctionListLibrary) FindMatches(argumentTypes ...string) []string {
	results := make([]string, 0, len(fl.funcsList))
	for _, fn := range fl.funcsList {
		if reflect.DeepEqual(fn.argumentTypes, argumentTypes) {
			results = append(results, fn.name)
			continue
		}

		if !fn.isVariadic || len(fn.argumentTypes) == 0 {
			continue
		}

		fnVariadicArgumentType := fn.argumentTypes[len(fn.argumentTypes)-1]
		index := len(argumentTypes) - 1
		for index >= 0 {
			if argumentTypes[index] != fnVariadicArgumentType {
				break
			}
			index--
		}

		// reach the end of arguments types
		if index == -1 && len(fn.argumentTypes) == 1 {
			results = append(results, fn.name)
			continue
		}

		remainingArgumentTypes := argumentTypes[:index+1] // +1 to include the index

		// remove from the fn args type any similar type from the right
		// before comparing remaining from both
		index = len(fn.argumentTypes) - 1
		for index >= 0 {
			if fn.argumentTypes[index] != fnVariadicArgumentType {
				break
			}
			index--
		}
		if reflect.DeepEqual(fn.argumentTypes[:index+1], remainingArgumentTypes) {
			results = append(results, fn.name)
		}
	}
	return results
}

type FunctionMapLibrary struct {
	funcsMap map[string][]Function
}

func (fm *FunctionMapLibrary) Register(functions ...Function) {
	for _, fn := range functions {
		key := strings.Join(fn.argumentTypes, ":")
		if _, exists := fm.funcsMap[key]; !exists {
			fm.funcsMap[key] = make([]Function, 0)
		}
		fm.funcsMap[key] = append(fm.funcsMap[key], fn)
	}
}

func (fm *FunctionMapLibrary) FindMatches(argumentTypes ...string) []string {
	return nil
}
