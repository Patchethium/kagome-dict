package data

import(
	"os"
	"time"
)


func dictUnidictCwBytes() ([]byte, error) {
	return _dictUnidictCw, nil
}

func dictUnidictCw() (*asset, error) {
	bytes, err := dictUnidictCwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
