package dependencyholder

import "go-structural-patterns/base"

// --- Pattern 2: Dependency Holder ---

// DependencyHolder はインターフェース型をフィールドとして持つ構造体
type DependencyHolder struct {
	A base.DoerA // インターフェース型をフィールドに持つ
	B base.DoerB
}

// NewDependencyHolder は DependencyHolder のインスタンスを生成する
// 依存性（インターフェースを満たす実装）を外部から注入する
func NewDependencyHolder(a base.DoerA, b base.DoerB) *DependencyHolder {
	return &DependencyHolder{
		A: a,
		B: b,
	}
}

// Note: DependencyHolder自体はDoerABインターフェースを実装しない
//       フィールドアクセスでメソッドを呼び出す: holder.A.DoA()
