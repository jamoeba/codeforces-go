package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func parseRawArray(rawArray string) (splits []string, err error) {
	invalidErr := fmt.Errorf("invalid test data: %s", rawArray)

	// check [] at leftmost and rightmost
	if len(rawArray) <= 1 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, invalidErr
	}

	// ignore [] at leftmost and rightmost
	rawArray = rawArray[1 : len(rawArray)-1]
	if rawArray == "" {
		return
	}

	isPoint := rawArray[0] == '('

	const sep = ','
	var depth, quotCnt, bracketCnt int
	for start := 0; start < len(rawArray); {
		end := start
	outer:
		for ; end < len(rawArray); end++ {
			switch rawArray[end] {
			case '[':
				depth++
			case ']':
				depth--
			case '"':
				quotCnt++
			case '(':
				bracketCnt++
			case ')':
				bracketCnt--
			case sep:
				if depth == 0 {
					if !isPoint {
						if quotCnt%2 == 0 {
							break outer
						}
					} else {
						if bracketCnt%2 == 0 {
							break outer
						}
					}
				}
			}
		}
		splits = append(splits, strings.TrimSpace(rawArray[start:end]))
		start = end + 1 // skip sep
	}
	if depth != 0 || quotCnt%2 != 0 {
		return nil, invalidErr
	}
	return
}

func parseRawArg(tp reflect.Type, rawData string) (v reflect.Value, err error) {
	rawData = strings.TrimSpace(rawData)
	invalidErr := fmt.Errorf("invalid test data: %s", rawData)
	switch tp.Kind() {
	case reflect.String:
		if len(rawData) <= 1 || rawData[0] != '"' || rawData[len(rawData)-1] != '"' {
			return reflect.Value{}, invalidErr
		}
		// remove " at leftmost and rightmost
		v = reflect.ValueOf(rawData[1 : len(rawData)-1])
	case reflect.Uint8: // byte
		// rawData like "a" or 'a'
		if len(rawData) != 3 || rawData[0] != '"' && rawData[0] != '\'' || rawData[2] != '"' && rawData[2] != '\'' {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(rawData[1])
	case reflect.Int:
		i, er := strconv.Atoi(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(i)
	case reflect.Uint:
		i, er := strconv.Atoi(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(uint(i))
	case reflect.Int64:
		i, er := strconv.Atoi(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(int64(i))
	case reflect.Float64:
		f, er := strconv.ParseFloat(rawData, 64)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(f)
	case reflect.Bool:
		b, er := strconv.ParseBool(rawData)
		if er != nil {
			return reflect.Value{}, invalidErr
		}
		v = reflect.ValueOf(b)
	case reflect.Slice:
		splits, er := parseRawArray(rawData)
		if er != nil {
			return reflect.Value{}, er
		}
		v = reflect.New(tp).Elem()
		for _, s := range splits {
			_v, er := parseRawArg(tp.Elem(), s)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.Append(v, _v)
		}
	case reflect.Ptr: // *TreeNode, *ListNode, *Point, *Interval
		switch tpName := tp.Elem().Name(); tpName {
		case "TreeNode":
			root, er := buildTreeNode(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(root)
		case "ListNode":
			head, er := buildListNode(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(head)
		case "Point": // nowcoder
			p, er := buildPoint(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(p)
		case "Interval": // nowcoder
			p, er := buildInterval(rawData)
			if er != nil {
				return reflect.Value{}, er
			}
			v = reflect.ValueOf(p)
		default:
			return reflect.Value{}, fmt.Errorf("unknown type %s", tpName)
		}
	default:
		return reflect.Value{}, fmt.Errorf("unknown type %s", tp.Name())
	}
	return
}

func toRawString(v reflect.Value) (s string, err error) {
	switch v.Kind() {
	case reflect.Slice:
		s = "["
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				s += ","
			}
			_s, er := toRawString(v.Index(i))
			if er != nil {
				return "", er
			}
			s += _s
		}
		s += "]"
	case reflect.Ptr: // *TreeNode, *ListNode, *Point, *Interval
		switch tpName := v.Type().Elem().Name(); tpName {
		case "TreeNode":
			s = v.Interface().(*TreeNode).toRawString()
		case "ListNode":
			s = v.Interface().(*ListNode).toRawString()
		case "Point":
			s = v.Interface().(*Point).toRawString()
		case "Interval":
			s = v.Interface().(*Interval).toRawString()
		default:
			return "", fmt.Errorf("unknown type %s", tpName)
		}
	case reflect.String:
		s = fmt.Sprintf(`"%s"`, v.Interface())
	case reflect.Uint8: // byte
		s = fmt.Sprintf(`"%c"`, v.Interface())
	case reflect.Float64:
		s = fmt.Sprintf(`%.5f`, v.Interface())
	default: // int uint int64 bool
		s = fmt.Sprintf(`%v`, v.Interface())
	}
	return
}

// rawExamples[i] = 输入+输出
func RunLeetCodeFuncWithExamples(t *testing.T, f interface{}, rawExamples [][]string, targetCaseNum int) (err error) {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return fmt.Errorf("f must be a function")
	}

	// 例如，-1 表示最后一个测试用例
	if targetCaseNum < 0 {
		targetCaseNum += len(rawExamples) + 1
	}

	allCasesOk := true
	fValue := reflect.ValueOf(f)
	for curCaseNum, example := range rawExamples {
		if targetCaseNum > 0 && curCaseNum+1 != targetCaseNum {
			continue
		}

		if len(example) != fType.NumIn()+fType.NumOut() {
			return fmt.Errorf("len(example) = %d, but we need %d+%d", len(example), fType.NumIn(), fType.NumOut())
		}

		rawIn := example[:fType.NumIn()]
		ins := make([]reflect.Value, len(rawIn))
		for i, rawArg := range rawIn {
			rawArg = trimSpaceAndNewLine(rawArg)
			ins[i], err = parseRawArg(fType.In(i), rawArg)
			if err != nil {
				return
			}
		}
		// just check rawExpectedOuts is valid or not
		rawExpectedOuts := example[fType.NumIn():]
		for i := range rawExpectedOuts {
			rawExpectedOuts[i] = trimSpaceAndNewLine(rawExpectedOuts[i])
			if _, err = parseRawArg(fType.Out(i), rawExpectedOuts[i]); err != nil {
				return
			}
		}

		const maxInputSize = 150
		inputInfo := strings.Join(rawIn, "\n")
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		outs := fValue.Call(ins)
		for i, out := range outs {
			rawActualOut, er := toRawString(out)
			if er != nil {
				return er
			}
			if !assert.Equal(t, rawExpectedOuts[i], rawActualOut, "Wrong Answer %d\nInput:\n%s", curCaseNum+1, inputInfo) {
				allCasesOk = false
			}
		}
	}

	if targetCaseNum > 0 && allCasesOk {
		t.Logf("case %d is ok", targetCaseNum)
		return RunLeetCodeFuncWithExamples(t, f, rawExamples, 0)
	}

	if allCasesOk {
		t.Log("OK")
	}

	return nil
}

func RunLeetCodeFuncWithCase(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string, targetCaseNum int) (err error) {
	examples := [][]string{}
	for i, input := range rawInputs {
		examples = append(examples, append(append([]string{}, input...), rawOutputs[i]...))
	}
	return RunLeetCodeFuncWithExamples(t, f, examples, targetCaseNum)
}

func RunLeetCodeFunc(t *testing.T, f interface{}, rawInputs [][]string, rawOutputs [][]string) error {
	return RunLeetCodeFuncWithCase(t, f, rawInputs, rawOutputs, 0)
}

// 方便打断点，配合 targetCaseNum 一起使用
var DebugCallIndex int

func RunLeetCodeClassWithExamples(t *testing.T, constructor interface{}, rawExamples [][3]string, targetCaseNum int) (err error) {
	cType := reflect.TypeOf(constructor)
	if cType.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a function")
	}
	if cType.NumOut() != 1 {
		return fmt.Errorf("constructor must have one and only one return value")
	}
	allCasesOk := true
	cFunc := reflect.ValueOf(constructor)

	// 例如，-1 表示最后一个测试用例
	if targetCaseNum < 0 {
		targetCaseNum += len(rawExamples) + 1
	}

	for curCase, example := range rawExamples {
		if targetCaseNum > 0 && curCase+1 != targetCaseNum {
			continue
		}

		names := strings.TrimSpace(example[0])
		inputArgs := strings.TrimSpace(example[1])
		rawExpectedOut := strings.TrimSpace(example[2])

		// parse called names
		// 调用 parseRawArray 确保数据是合法的
		methodNames, er := parseRawArray(names)
		if er != nil {
			return er
		}
		for i, name := range methodNames {
			name = name[1 : len(name)-1] // 移除引号
			name = strings.Title(name)   // 首字母大写以匹配模板方法名称
			methodNames[i] = name
		}

		// parse inputs
		rawArgsList, er := parseRawArray(inputArgs)
		if er != nil {
			return er
		}
		if len(rawArgsList) != len(methodNames) {
			return fmt.Errorf("invalid test data: mismatch names and input args (%d != %d)", len(methodNames), len(rawArgsList))
		}

		// parse constructor input
		constructorArgs, er := parseRawArray(rawArgsList[0])
		if er != nil {
			return er
		}
		constructorIns := make([]reflect.Value, len(constructorArgs))
		for i, arg := range constructorArgs {
			constructorIns[i], err = parseRawArg(cType.In(i), arg)
			if err != nil {
				return
			}
		}

		// call constructor, get struct instance
		obj := cFunc.Call(constructorIns)[0]

		// use a pointer to call methods
		pObj := reflect.New(obj.Type())
		pObj.Elem().Set(obj)

		if DebugCallIndex < 0 {
			DebugCallIndex += len(rawArgsList)
		}
		rawActualOut := "[null"
		for callIndex := 1; callIndex < len(rawArgsList); callIndex++ {
			name := methodNames[callIndex]
			method := pObj.MethodByName(name)
			emptyValue := reflect.Value{}
			if method == emptyValue {
				return fmt.Errorf("invalid test data: %s", methodNames[callIndex])
			}
			methodType := method.Type()

			// parse method input
			methodArgs, er := parseRawArray(rawArgsList[callIndex])
			if er != nil {
				return er
			}
			in := make([]reflect.Value, methodType.NumIn()) // 注意：若入参为空，methodArgs 可能是 [] 也可能是 [null]
			for i := range in {
				in[i], err = parseRawArg(methodType.In(i), methodArgs[i])
				if err != nil {
					return
				}
			}

			if callIndex == DebugCallIndex {
				print()
			}
			// call method
			if actualOuts := method.Call(in); len(actualOuts) > 0 {
				s, er := toRawString(actualOuts[0])
				if er != nil {
					return er
				}
				rawActualOut += "," + s
			} else {
				rawActualOut += ",null"
			}
		}
		rawActualOut += "]"

		// 比较前，去除 rawExpectedOut 中逗号后的空格
		// todo: 提示错在哪个 callIndex 上
		rawExpectedOut = strings.ReplaceAll(rawExpectedOut, ", ", ",")
		if !assert.Equal(t, rawExpectedOut, rawActualOut, "Wrong Answer %d", curCase+1) {
			allCasesOk = false
		}
	}

	if targetCaseNum > 0 && allCasesOk {
		t.Logf("case %d is ok", targetCaseNum)
		return RunLeetCodeClassWithExamples(t, constructor, rawExamples, 0)
	}

	if allCasesOk {
		t.Log("OK")
	}

	return nil
}

func RunLeetCodeClassWithCase(t *testing.T, constructor interface{}, rawInputs, rawOutputs []string, targetCaseNum int) (err error) {
	examples := [][3]string{}
	for i, input := range rawInputs {
		input = strings.TrimSpace(input)
		lines := strings.Split(input, "\n")
		examples = append(examples, [3]string{lines[0], lines[1], rawOutputs[i]})
	}
	return RunLeetCodeClassWithExamples(t, constructor, examples, targetCaseNum)
}

func RunLeetCodeClass(t *testing.T, constructor interface{}, rawInputs, rawOutputs []string) error {
	return RunLeetCodeClassWithCase(t, constructor, rawInputs, rawOutputs, 0)
}

// 无尽对拍模式
func CompareInf(t *testing.T, inputGenerator, runACFunc, runFunc interface{}) {
	const needPrint = runtime.GOOS == "darwin"

	ig := reflect.ValueOf(inputGenerator)
	if ig.Kind() != reflect.Func {
		t.Fatal("input generator must be a function")
	}
	runAC := reflect.ValueOf(runACFunc)
	run := reflect.ValueOf(runFunc)
	// just check numbers
	if !assert.Equal(t, run.Type().NumIn(), runAC.Type().NumIn()) ||
		!assert.Equal(t, run.Type().NumOut(), runAC.Type().NumOut()) {
		t.Fatal("different input/output")
	}

	for tc := 1; ; tc++ {
		inArgs := ig.Call(nil)

		// 先生成字符串，以免 inArgs 被修改
		insStr := []byte{}
		for i, arg := range inArgs {
			if i > 0 {
				insStr = append(insStr, '\n')
			}
			s, err := toRawString(arg)
			if err != nil {
				t.Fatal(err)
			}
			insStr = append(insStr, s...)
		}

		// todo deep copy slice
		expectedOut := runAC.Call(inArgs)
		actualOut := run.Call(inArgs)

		for i, eOut := range expectedOut {
			if !assert.Equal(t, eOut.Interface(), actualOut[i].Interface(), "Wrong Answer %d\nInput:\n%s", tc, insStr) && needPrint {
				fmt.Printf("[CASE %d]\n", tc)
				fmt.Println("[AC]", eOut.Interface())
				fmt.Println("[WA]", actualOut[i].Interface())
				fmt.Printf("[INPUT]\n%s\n\n", insStr)
			}
		}

		if tc%1e5 == 0 {
			s := fmt.Sprintf("%d cases passed.", tc)
			t.Log(s)
			if needPrint {
				fmt.Println(s)
			}
		}
	}
}
