package base

// --- Existing Interfaces ---
// DoerA はAの操作を行うインターフェース
type DoerA interface {
	DoA() string
}

// DoerB はBの操作を行うインターフェース
type DoerB interface {
	DoB() string
}

// DoerAB はAとB両方の操作を行うインターフェース
type DoerAB interface {
	DoerA // DoerA のメソッドセットを含む
	DoerB // DoerB のメソッドセットを含む
}

// --- New Interfaces Demonstrating Inclusion ---

// DoerC は DoerA の仕様を含み、追加のメソッド DoC を持つインターフェース
type DoerC interface {
	DoerA // DoerA のメソッドセットを含む
	DoC() string
}

// DoerD は DoerB の仕様を含み、追加のメソッド DoD を持つインターフェース
type DoerD interface {
	DoerB // DoerB のメソッドセットを含む
	DoD() string
}

// DoerCD は DoerC と DoerD の仕様を両方含むインターフェース
// 結果として DoerA と DoerB のメソッドセットも含むことになる
type DoerCD interface {
	DoerC // DoerC のメソッドセットを含む (DoA, DoC)
	DoerD // DoerD のメソッドセットを含む (DoB, DoD)
}
