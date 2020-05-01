package chapter3

// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

func (k1 Kadai1) ID() int {
	return k1.id
}

func (k1 *Kadai1) SetID(id int) {
	k1.id = id
}

func (k1 Kadai1) Name() string {
	return k1.name
}

func (k1 *Kadai1) SetName(name string) {
	k1.name = name
}
