// Code generated by "stringer -type SelectionType selection.go"; DO NOT EDIT

package prompt

import "fmt"

const _SelectionType_name = "CHARACTERSLINESBLOCK"

var _SelectionType_index = [...]uint8{0, 10, 15, 20}

func (i SelectionType) String() string {
	if i < 0 || i >= SelectionType(len(_SelectionType_index)-1) {
		return fmt.Sprintf("SelectionType(%d)", i)
	}
	return _SelectionType_name[_SelectionType_index[i]:_SelectionType_index[i+1]]
}