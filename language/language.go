package language

import (
	"./assembly"
	"./ats"
	"./bash"
	"./c"
	"./clojure"
	"./cobol"
	"./coffeescript"
	"./cpp"
	"./crystal"
	"./csharp"
	"./d"
	"./elixir"
	"./elm"
	"./erlang"
	"./fsharp"
	"./golang"
	"./groovy"
	"./haskell"
	"./idris"
	"./java"
	"./javascript"
	"./julia"
	"./kotlin"
	"./lua"
	"./mercury"
	"./nim"
	"./ocaml"
	"./perl"
	"./perl6"
	"./php"
	"./python"
	"./ruby"
	"./rust"
	"./scala"
	"./swift"
	"./typescript"
)

type runFn func([]string, int64, string) (string, string, error, int64, int64)

var languages = map[string]runFn{
	"assembly":     assembly.Run,
	"ats":          ats.Run,
	"bash":         bash.Run,
	"c":            c.Run,
	"clojure":      clojure.Run,
	"cobol":        cobol.Run,
	"coffeescript": coffeescript.Run,
	"crystal":      crystal.Run,
	"csharp":       csharp.Run,
	"d":            d.Run,
	"elixir":       elixir.Run,
	"elm":          elm.Run,
	"cpp":          cpp.Run,
	"erlang":       erlang.Run,
	"fsharp":       fsharp.Run,
	"haskell":      haskell.Run,
	"idris":        idris.Run,
	"go":           golang.Run,
	"groovy":       groovy.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"julia":        julia.Run,
	"kotlin":       kotlin.Run,
	"lua":          lua.Run,
	"mercury":      mercury.Run,
	"nim":          nim.Run,
	"ocaml":        ocaml.Run,
	"perl":         perl.Run,
	"perl6":        perl6.Run,
	"php":          php.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"rust":         rust.Run,
	"scala":        scala.Run,
	"swift":        swift.Run,
	"typescript":   typescript.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, maxTimeout int64, files []string, stdin string) (string, string, error, int64, int64) {
	return languages[lang](files, maxTimeout, stdin)
}
