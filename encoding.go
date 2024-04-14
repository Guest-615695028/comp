package comp
import "strconv"
// Return the JSON for (error if any, value if none)
func ValueJSON(v any, err error) string {
	if err != nil {
		return QuoteToJSON(err.Error())
	}
	switch w := v.(type) {
	case nil:
		return "null"
	case bool:
		return Cond(w, "true", "false")
	case int64:
		return strconv.FormatInt(w, 10)
	case int32:
		return strconv.FormatInt(int64(w), 10)
	case int16:
		return strconv.FormatInt(int64(w), 10)
	case int8:
		return strconv.FormatInt(int64(w), 10)
	case int:
		return strconv.FormatInt(int64(w), 10)
	case uint64:
		return strconv.FormatUint(w, 10)
	case uint32:
		return strconv.FormatUint(uint64(w), 10)
	case uint16:
		return strconv.FormatUint(uint64(w), 10)
	case uint8:
		return strconv.FormatUint(uint64(w), 10)
	case uint:
		return strconv.FormatUint(uint64(w), 10)
	case uintptr:
		return strconv.FormatUint(uint64(w), 10)
	case float32:
		s := strconv.FormatFloat(float64(w), 'g', -1, 32)
		switch s {
		case "NaN", "-Inf", "+Inf":
			return "\"" + s + "\""
		default:
			return s
		}
	case float64:
		s := strconv.FormatFloat(w, 'g', -1, 64)
		switch s {
		case "NaN", "-Inf", "+Inf":
			return "\"" + s + "\""
		default:
			return s
		}
	case complex64:
		s := strconv.FormatComplex(complex128(w), 'g', -1, 32)
		return "\"" + s[1:len(s)-1] + "\""
	case complex128:
		s := strconv.FormatComplex(w, 'g', -1, 32)
		return "\"" + s[1:len(s)-1] + "\""
	case string:
		return QuoteToJSON(w)
	case []rune:
		return QuoteToJSON(string(w))
	case []byte:
		return QuoteToJSON(string(w))
	}
	return "{}" //Temporarily reserved empty object
}

// Last Hexadigit
func Hex(i byte, cap bool) byte {
	if cap {
		return "0123456789ABCDEF"[i&15]
	} else {
		return "0123456789abcdef"[i&15]
	}
}
func Escape(r rune, cap bool) []byte {
	s := [6]byte{'\\', 'u'}
	for i := 3; i >= 0; i-- {
		s[5-i] = Hex(byte(r>>i), cap)
	}
	return s[:]
}

// Quote friendly to JSON string
func QuoteToJSON(s string) string {
	q := []byte{}
	for i, v := range s {
		switch v {
		case '\b':
			q = append(q, '\\', 'b')
		case '\f':
			q = append(q, '\\', 'f')
		case '\n':
			q = append(q, '\\', 'n')
		case '\r':
			q = append(q, '\\', 'r')
		case '\t':
			q = append(q, '\\', 't')
		case '"':
			q = append(q, '\\', '"')
		case '\\':
			q = append(q, '\\', '\\')
		case '/':
			q = append(q, '\\', '/')
		case 0xFFFD:
			q = append(append(q, "\\ufffd(0x"...),
				Hex(s[i]>>4, false), Hex(s[i], false), ')')
		default:
			if v >= 0x20 && v <= 0x7E {
				q = append(q, byte(v))
			} else if v >= 0x10000 {
				v -= 0x10000
				h, l := 0xD800+(v>>10&0x03FF), 0xDC00+(v&0x03FF)
				q = append(append(q, Escape(h, false)...), Escape(l, false)...)
			} else {
				q = append(q, Escape(v, false)...)
			}
		}
	}
	return string(q)
}
