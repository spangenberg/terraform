// Code generated by "stringer -type=TokenType"; DO NOT EDIT

package scanner

import "fmt"

const _TokenType_name = "BEGINPERCENTOPARENCPARENSTARPLUSCOMMAMINUSPERIODSLASHBOOLFLOATINTEGERSTRINGOBRACKETCBRACKETIDENTIFIERLITERALENDOQUOTECQUOTEEOFINVALID"

var _TokenType_map = map[TokenType]string{
	36:    _TokenType_name[0:5],
	37:    _TokenType_name[5:12],
	40:    _TokenType_name[12:18],
	41:    _TokenType_name[18:24],
	42:    _TokenType_name[24:28],
	43:    _TokenType_name[28:32],
	44:    _TokenType_name[32:37],
	45:    _TokenType_name[37:42],
	46:    _TokenType_name[42:48],
	47:    _TokenType_name[48:53],
	66:    _TokenType_name[53:57],
	70:    _TokenType_name[57:62],
	73:    _TokenType_name[62:69],
	83:    _TokenType_name[69:75],
	91:    _TokenType_name[75:83],
	93:    _TokenType_name[83:91],
	105:   _TokenType_name[91:101],
	111:   _TokenType_name[101:108],
	125:   _TokenType_name[108:111],
	8220:  _TokenType_name[111:117],
	8221:  _TokenType_name[117:123],
	9220:  _TokenType_name[123:126],
	65533: _TokenType_name[126:133],
}

func (i TokenType) String() string {
	if str, ok := _TokenType_map[i]; ok {
		return str
	}
	return fmt.Sprintf("TokenType(%d)", i)
}
