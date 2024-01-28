package pkg

import "os"

const StorageName string = "GoFileVault"

func InitLocalDirectory() {
	os.Mkdir(StorageName, 0755)
}
