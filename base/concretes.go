package base

// --- Existing Concrete Implementations ---
// ConcreteDoerA は DoerA の具体的な実装
type ConcreteDoerA struct{}

func (d *ConcreteDoerA) DoA() string { return "Did A concretely" }

// ConcreteDoerB は DoerB の具体的な実装
type ConcreteDoerB struct{}

func (d *ConcreteDoerB) DoB() string { return "Did B concretely" }

// --- New Concrete Implementation for DoerCD ---

// ConcreteDoerCD は DoerCD インターフェースの具体的な実装
// DoerCD は DoerA, DoerB, DoerC, DoerD のメソッドをすべて要求するため、
// この構造体はそれらすべてを実装する必要がある
type ConcreteDoerCD struct{}

// DoerA のメソッド
func (d *ConcreteDoerCD) DoA() string { return "ConcreteCD doing A" }

// DoerB のメソッド
func (d *ConcreteDoerCD) DoB() string { return "ConcreteCD doing B" }

// DoerC のメソッド
func (d *ConcreteDoerCD) DoC() string { return "ConcreteCD doing C" }

// DoerD のメソッド
func (d *ConcreteDoerCD) DoD() string { return "ConcreteCD doing D" }

// コンパイル時に ConcreteDoerCD が各インターフェースを満たすかチェック (任意)
var _ DoerA = (*ConcreteDoerCD)(nil)
var _ DoerB = (*ConcreteDoerCD)(nil)
var _ DoerC = (*ConcreteDoerCD)(nil)
var _ DoerD = (*ConcreteDoerCD)(nil)
var _ DoerCD = (*ConcreteDoerCD)(nil)

// ★ DoerCD を満たすものは、DoerAB のメソッドセットも持っているため、
//   DoerAB も満たすことを確認
var _ DoerAB = (*ConcreteDoerCD)(nil)
