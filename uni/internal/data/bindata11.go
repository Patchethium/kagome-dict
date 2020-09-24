package data

import(
	"os"
	"time"
)


func dictUnidictAlBytes() ([]byte, error) {
	return _dictUnidictAl, nil
}

func dictUnidictAl() (*asset, error) {
	bytes, err := dictUnidictAlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.al", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
