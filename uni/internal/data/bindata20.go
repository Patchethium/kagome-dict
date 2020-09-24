package data

import(
	"os"
	"time"
)


func dictUnidictAuBytes() ([]byte, error) {
	return _dictUnidictAu, nil
}

func dictUnidictAu() (*asset, error) {
	bytes, err := dictUnidictAuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.au", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
