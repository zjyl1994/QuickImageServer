package main

import(
	"github.com/seehuhn/mt19937"
	"time"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"fmt"
	"os"
)

func ImageID2Path(imageid string)string{
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s.jpg",conf.Storage,imageid[0:2],imageid[2:4],imageid[4:6],imageid[6:8],imageid[8:10],imageid[10:12],imageid[12:14],imageid[14:16])
}

func MakeImageID()string{
	mt:=mt19937.New()
	mt.Seed(time.Now().UnixNano())
	var buf = make([]byte, 8)
    binary.BigEndian.PutUint64(buf, mt.Uint64())
    return strings.ToUpper(hex.EncodeToString(buf))
}

func FileExist(filename string)bool{
	if _, err := os.Stat(filename); err != nil {
		return false
	}else{
		return true
	}
}

func BuildTree(imageid string)error{
	return os.MkdirAll(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s",conf.Storage,imageid[0:2],imageid[2:4],imageid[4:6],imageid[6:8],imageid[8:10],imageid[10:12],imageid[12:14]),0666)
}