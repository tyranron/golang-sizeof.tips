package app

import (
	"fmt"
	. "go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"unsafe"
)

const exampleCode = `
struct {
	a string
	b bool
	c string
}
`

var basicSizes = map[string]uint64{
	"bool":       uint64(unsafe.Sizeof(true)),
	"int8":       uint64(unsafe.Sizeof(int8(0))),
	"uint8":      uint64(unsafe.Sizeof(uint8(0))),
	"byte":       uint64(unsafe.Sizeof(byte(0))),
	"int16":      uint64(unsafe.Sizeof(int16(0))),
	"uint16":     uint64(unsafe.Sizeof(uint16(0))),
	"int32":      uint64(unsafe.Sizeof(int32(0))),
	"uint32":     uint64(unsafe.Sizeof(uint32(0))),
	"rune":       uint64(unsafe.Sizeof(rune(0))),
	"float32":    uint64(unsafe.Sizeof(float32(0))),
	"int":        uint64(unsafe.Sizeof(int(0))),
	"uint":       uint64(unsafe.Sizeof(uint(0))),
	"int64":      uint64(unsafe.Sizeof(int64(0))),
	"uint64":     uint64(unsafe.Sizeof(uint64(0))),
	"float64":    uint64(unsafe.Sizeof(float64(0))),
	"uintptr":    uint64(unsafe.Sizeof(uintptr(0))),
	"complex64":  uint64(unsafe.Sizeof(complex64(0))),
	"complex128": uint64(unsafe.Sizeof(complex128(0))),
	"string":     uint64(unsafe.Sizeof("")),
}

var fixedSizes = map[string]uint64{
	"ptr":   uint64(unsafe.Sizeof(&struct{}{})),
	"map":   uint64(unsafe.Sizeof(map[bool]bool{})),
	"slice": uint64(unsafe.Sizeof([]struct{}{})),
	"chan":  uint64(unsafe.Sizeof(make(chan struct{}))),
	"func":  uint64(unsafe.Sizeof(func() {})),
}

func discoverCode(code string) (string, error) {
	expr, err := parser.ParseExpr(code)
	if err != nil {
		return "", err
	}

	size, err := discover(expr)
	if err != nil {
		return "", err
	}
	//vis := &myVis{}
	//Walk(vis, expr)
	return fmt.Sprintf("%+v\n%s\n--%d\n", expr, reflect.TypeOf(expr), size), nil
}

type myVis struct {
	s string
}

func (v *myVis) Visit(node Node) Visitor {
	v.s += fmt.Sprintf("%+v\n|||\n", node)
	return v
}

func discover(node Node) (uint64, error) {
	switch n := node.(type) {
	case *Ident:
		size, exists := basicSizes[n.Name]
		if !exists {
			return 0, fmt.Errorf("unknown type '%s'", n.Name)
		}
		return size, nil
	case *StarExpr:
		return fixedSizes["ptr"], nil
	case *MapType:
		return fixedSizes["map"], nil
	case *ChanType:
		return fixedSizes["chan"], nil
	case *FuncLit:
		return fixedSizes["func"], nil
	case *ArrayType:
		if n.Len == nil {
			return fixedSizes["slice"], nil
		}
		len, ok := n.Len.(*BasicLit)
		if !ok || len.Kind != token.INT {
			return 0, fmt.Errorf("invalid length in array definition")
		}
		num, err := strconv.ParseUint(len.Value, 10, 64)
		switch {
		case err != nil:
			return 0, fmt.Errorf("invalid length in array definition")
		case num < 1:
			return 0, nil
		}
		size, err := discover(n.Elt)
		if err != nil {
			return 0, err
		}
		return size * uint64(num), nil
	case *StructType:
		if len(n.Fields.List) < 1 {
			return 0, nil
		}
		var err error
		alignment := uint64(0)
		sizes := make([]uint64, len(n.Fields.List))
		for i, field := range n.Fields.List {
			sizes[i], err = discover(field.Type)
			if err != nil {
				return 0, err
			}
			if sizes[i] > alignment {
				alignment = sizes[i]
			}
		}
		if alignment > basicSizes["uintptr"] {
			alignment = basicSizes["uintptr"]
		}
		num, size := uint64(0), uint64(0)
		for _, s := range sizes {
			if s == 0 {
				continue
			}
			size += s
			if size < alignment {
				continue
			}
			if size-s > 0 {
				num++
			}
			num += s / alignment
			size = s % alignment
		}
		if size > 0 {
			num++
		}
		return num * alignment, nil
	default:
		return 0, fmt.Errorf("invalid type expression")
	}
}
