package delegator

import "go-structural-patterns/base"

// --- Pattern 4: Delegator ---

// Delegator は単一のインターフェースに処理を委譲する構造体
type Delegator struct {
	delegate base.DoerAB // 委譲先のインターフェース
}

// NewDelegator は Delegator のインスタンスを生成する
func NewDelegator(ab base.DoerAB) *Delegator {
	return &Delegator{delegate: ab}
}

// Delegator が DoerA インターフェースを実装 (委譲)
func (d *Delegator) DoA() string {
	return d.delegate.DoA() // 処理を内部のdelegateに委譲
}

// Delegator が DoerB インターフェースを実装 (委譲)
func (d *Delegator) DoB() string {
	return d.delegate.DoB() // 処理を内部のdelegateに委譲
}

// コンパイル時に DoerAB を満たしているかチェック (任意)
var _ base.DoerAB = (*Delegator)(nil)
