package data

import(
	"os"
	"time"
)


func dictUnidictAyBytes() ([]byte, error) {
	return _dictUnidictAy, nil
}

func dictUnidictAy() (*asset, error) {
	bytes, err := dictUnidictAyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.ay", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
