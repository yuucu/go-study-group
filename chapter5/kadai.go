package chapter5

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type cutEntity struct {
	args      []string
	delimiter string
	fields    int
}

var delimiter = flag.String("d", ",", "区切り文字を指定してください")
var fields = flag.Int("f", 1, "フィールドの何番目を取り出すか指定してください")

// go-cutコマンドを実装しよう
func main() {
	flag.Parse()

	cutEntity := cutEntity{
		args:      flag.Args(),
		delimiter: *delimiter,
		fields:    *fields,
	}

	err := Validate(cutEntity)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(cutEntity.args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = Cut(cutEntity, file, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
func Cut(cutEntity cutEntity, r io.Reader, w io.Writer) error {

	// この部分をCutコマンドとして関数2つ目に切り出す------
	// ヒント：NewScannerにfileを渡しているが、NewScannerはio.Readerであれば何でも良い
	// また、出力も現在fmt.Println(s)にしているが、io.Writerを使って書き出す先を指定できるようにしてやる
	// 関数の引数で読み出すio.Readerと、
	// 書き出すio.Writer (本関数からはos.Stdout, テストからはbyte.Bufferなどへ)を指定できるようにすると良い
	bw := bufio.NewWriter(w)
	defer bw.Flush()

	scanner := bufio.NewScanner(r)
	delimiter := []byte(cutEntity.delimiter)
	for scanner.Scan() {
		btext := scanner.Bytes()
		sb := bytes.Split(btext, delimiter)
		if len(sb) < cutEntity.fields {
			return fmt.Errorf("-fの値に該当するデータがありません")
		}
		b := sb[cutEntity.fields-1]
		_, err := bw.Write(b)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
	// ------------------------------------------------
}

func Validate(cutEntity cutEntity) error {
	// このValidationを関数1つ目に切り出す ---------
	// ヒント：flagの内容を渡してやって、バリデーションし、エラーがあれば返すような関数にできる
	if len(cutEntity.args) == 0 {
		return fmt.Errorf("ファイルパスを指定してください。")
	}
	if cutEntity.fields < 0 {
		return fmt.Errorf("-f は1以上である必要があります")
	}
	return nil
	// ---------------------------------------
}
