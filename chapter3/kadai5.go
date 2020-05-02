package chapter3

import "encoding/json"

// 課題5
// とあるマスターデータのテーブルの構造をモデル化したstructを作りました。
// idやnameを変更できないようにカプセル化し、Getterのみ実装しました。
// しかしモデルをjson.Marshal()でシリアライズすると{}となってしまい、id, nameが出力されません。
// １つ関数を実装してシリアライズした結果が次のようになるようにしましょう。
// {"id": 1, "name": "hoge"}
// ヒント https://golang.org/pkg/encoding/json/#Marshal
// >> If an encountered value implements the Marshaler interface and is not a nil pointer, Marshal calls its MarshalJSON method to produce JSON
type Master struct {
	id   int
	name string
}

func (m Master) ID() int {
	return m.id
}

func (m Master) Name() string {
	return m.name
}

func (m Master) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   m.id,
		Name: m.name,
	})

	if err != nil {
		return nil, err
	}
	return j, nil
}
