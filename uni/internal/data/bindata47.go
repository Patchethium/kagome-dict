package data

import(
	"os"
	"time"
)


func dictUnidictBvBytes() ([]byte, error) {
	return _dictUnidictBv, nil
}

func dictUnidictBv() (*asset, error) {
	bytes, err := dictUnidictBvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
