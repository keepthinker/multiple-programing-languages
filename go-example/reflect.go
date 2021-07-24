package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type staff struct {
	name string `json:"name"`
	age  int32  `json:"age"`
}

func (s *staff) SpeakName() string {
	fmt.Println("My name is", s.name)
	return s.name
}

func (s *staff) getAge() int32 {
	return s.age
}

func (s *staff) UpdateName(name string) {
	s.name = name
}

// main most codes from book namedThe Go Rrogramming Language
func main() {
	fmt.Println("reflect example")
	s := staff{name: "ken", age: 12}
	fmt.Println(s.name, ":", s.getAge())
	printType(s)

	fmt.Println("convert to string")
	fmt.Println(stringValue(1.2329809089793))
	fmt.Println(stringValue(100))
	fmt.Println(stringValue(true))

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		}, Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("strangelove", strangelove)

	fmt.Println("------------------")
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()
	fmt.Println("a CanAddr()", a.CanAddr())
	fmt.Println("b CanAddr()", b.CanAddr())
	fmt.Println("c CanAddr()", c.CanAddr())
	fmt.Println("d CanAddr()", d.CanAddr())

	dPointer := d.Addr().Interface().(*int)
	*dPointer = 3
	fmt.Println("d after update 3:", d)

	d.SetInt(10)
	fmt.Println("d after update 10:", d)

	d.Set(reflect.ValueOf(20))
	fmt.Println("d after update 20:", d)

	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println("stdout type:", stdout.Type())
	pdf := stdout.FieldByName("pfd")
	fmt.Println("pfd:", pdf)
	fmt.Println("can addr:", pdf.CanAddr(), " can set", pdf.CanSet())

	printFieldTags(&s)

	fmt.Println("-----------------print method---------------------")
	printMethods(time.Hour)

	callMethod(&s)
}

func callMethod(s interface{}) {
	v := reflect.ValueOf(s)
	msn := v.MethodByName("SpeakName")
	if msn.IsValid() {
		r := msn.Call(nil)
		for i := 0; i < len(r); i++ {
			fmt.Println(i, ":", r[i])
		}
	}

	mun := v.MethodByName("UpdateName")
	if mun.IsValid() {
		in := []reflect.Value{reflect.ValueOf("Michael")}
		mun.Call(in)
	}

	if msn.IsValid() {
		r := msn.Call(nil)
		for i := 0; i < len(r); i++ {
			fmt.Println(i, ":", r[i])
		}
	}

}

func printMethods(x interface{}) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Printf("type %s\n", t)
	for i := 0; i < t.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}

}

func printFieldTags(ptr interface{}) {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		fmt.Println("tag:", tag)
		fmt.Println("tag.key:", tag.Get("json"))
		fmt.Println("field.value:", v.Field(i))

	}
}

func printType(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("type string:", t.String())
	fmt.Println("type: ", t)

	v := reflect.ValueOf(i)
	fmt.Println("value:", v)
	fmt.Printf("value: %v\n", v)
	fmt.Println("value: ", v.String())

	fmt.Println("type: ", v.Type().String())

	tn := reflect.New(t)

	m := tn.MethodByName("SpeakName")
	if m.IsValid() {
		m.Call(nil)
	} else {
		fmt.Println("failt to get MethodByName")
	}

}

func stringValue(i interface{}) string {
	ivalue := reflect.ValueOf(i)

	switch ivalue.Kind() {
	case reflect.String:
		return ivalue.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(ivalue.Int(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(ivalue.Float(), 'e', 12, 32)
	case reflect.Float64:
		return strconv.FormatFloat(ivalue.Float(), 'e', 12, 64)
	case reflect.Bool:
		return strconv.FormatBool(ivalue.Bool())
	default:
		return ivalue.Type().String() + " value"
	}
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
