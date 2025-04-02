package compositedoer

import "go-structural-patterns/base"

// --- Pattern 3: Composite Implementer ---

// CompositeDoer は複数のインターフェースを保持し、自身もインターフェースを実装する構造体
type CompositeDoer struct {
	delegateA base.DoerA // 内部で保持するインターフェース (非公開でも良い)
	delegateB base.DoerB
}

// NewCompositeDoer は CompositeDoer のインスタンスを生成する
func NewCompositeDoer(a base.DoerA, b base.DoerB) *CompositeDoer {
	return &CompositeDoer{
		delegateA: a,
		delegateB: b,
	}
}

// CompositeDoer が DoerA インターフェースを実装 (委譲)
func (c *CompositeDoer) DoA() string {
	return c.delegateA.DoA() // 処理を内部のdelegateAに委譲
}

// CompositeDoer が DoerB インターフェースを実装 (委譲)
func (c *CompositeDoer) DoB() string {
	return c.delegateB.DoB() // 処理を内部のdelegateBに委譲
}

// コンパイル時に DoerAB を満たしているかチェック (任意)
var _ base.DoerAB = (*CompositeDoer)(nil)
