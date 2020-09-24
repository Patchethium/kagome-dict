package data

import(
	"os"
	"time"
)


func dictUnidictBpBytes() ([]byte, error) {
	return _dictUnidictBp, nil
}

func dictUnidictBp() (*asset, error) {
	bytes, err := dictUnidictBpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
