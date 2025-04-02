package main

import (
	"fmt"
	"go-structural-patterns/base"
	"go-structural-patterns/compositedoer"
	"go-structural-patterns/concreteholder"
	"go-structural-patterns/delegator"
	"go-structural-patterns/dependencyholder"
)

func main() {
	// 共通の実装を作成
	concreteA := &base.ConcreteDoerA{}
	concreteB := &base.ConcreteDoerB{}

	fmt.Println("--- Pattern 1: Concrete Holder ---")
	holder1 := concreteholder.NewConcreteHolder()
	fmt.Println("holder1.A.DoA():", holder1.A.DoA()) // フィールド経由で呼び出し
	fmt.Println("holder1.B.DoB():", holder1.B.DoB()) // フィールド経由で呼び出し

	fmt.Println("\n--- Pattern 2: Dependency Holder ---")
	holder2 := dependencyholder.NewDependencyHolder(concreteA, concreteB)
	fmt.Println("holder2.A.DoA():", holder2.A.DoA()) // フィールド経由で呼び出し
	fmt.Println("holder2.B.DoB():", holder2.B.DoB()) // フィールド経由で呼び出し

	fmt.Println("\n--- Pattern 3: Composite Implementer ---")
	// CompositeDoer は DoerAB を実装する
	var doerAB3 base.DoerAB = compositedoer.NewCompositeDoer(concreteA, concreteB)
	fmt.Println("doerAB3.DoA():", doerAB3.DoA()) // インターフェース経由で呼び出し
	fmt.Println("doerAB3.DoB():", doerAB3.DoB()) // インターフェース経由で呼び出し

	fmt.Println("\n--- Pattern 4: Delegator ---")
	// Delegator に渡すための DoerAB が必要 (例えば CompositeDoer を使う)
	underlyingDoer := compositedoer.NewCompositeDoer(concreteA, concreteB)
	var doerAB4 base.DoerAB = delegator.NewDelegator(underlyingDoer)
	fmt.Println("doerAB4.DoA():", doerAB4.DoA()) // インターフェース経由で呼び出し
	fmt.Println("doerAB4.DoB():", doerAB4.DoB()) // インターフェース経由で呼び出し

	fmt.Println("\n--- Interface Inclusion Verification ---")
	// ConcreteDoerCD のインスタンスを作成
	concreteCD := &base.ConcreteDoerCD{}

	// 1. concreteCD は DoerCD インターフェースを満たす
	var iCD base.DoerCD = concreteCD
	fmt.Println("iCD.DoA():", iCD.DoA())
	fmt.Println("iCD.DoB():", iCD.DoB())
	fmt.Println("iCD.DoC():", iCD.DoC())
	fmt.Println("iCD.DoD():", iCD.DoD())

	// 2. ★★★ ここがポイント ★★★
	// DoerCD は DoerC と DoerD を含み、
	// DoerC は DoerA を、DoerD は DoerB を含むため、
	// DoerCD は結果的に DoerA と DoerB のメソッドセットを持つことになる。
	// したがって、DoerCD を満たす型は、DoerAB インターフェースも満たす。
	var iAB base.DoerAB = concreteCD // エラーなく代入可能！
	fmt.Println("iAB assigned from concreteCD successfully!")
	fmt.Println("iAB.DoA():", iAB.DoA())
	fmt.Println("iAB.DoB():", iAB.DoB())

	// 参考: もちろん、個別のインターフェースも満たしている
	var iA base.DoerA = concreteCD
	var iB base.DoerB = concreteCD
	var iC base.DoerC = concreteCD
	var iD base.DoerD = concreteCD
	fmt.Println("Also satisfies DoerA, DoerB, DoerC, DoerD:",
		iA.DoA() != "", iB.DoB() != "", iC.DoC() != "", iD.DoD() != "")

}
