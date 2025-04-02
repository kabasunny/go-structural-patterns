package concreteholder

import "go-structural-patterns/base"

// --- Pattern 1: Concrete Holder ---

// ConcreteHolder は具象型を直接保持する構造体
type ConcreteHolder struct {
	A base.ConcreteDoerA // 具体的な実装をフィールドに持つ
	B base.ConcreteDoerB
}

// NewConcreteHolder は ConcreteHolder のインスタンスを生成する
func NewConcreteHolder() *ConcreteHolder {
	return &ConcreteHolder{
		A: base.ConcreteDoerA{},
		B: base.ConcreteDoerB{},
	}
}

// Note: ConcreteHolder自体はDoerABインターフェースを実装しない
//       フィールドアクセスでメソッドを呼び出す: holder.A.DoA()
