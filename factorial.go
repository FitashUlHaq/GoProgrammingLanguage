//Package factorial calculates the factorial
package factorial

//initialfactorial is a local variable
var initialfactorial int = 1

//Calculatefactorial is a global function
func Calculatefactorial(count int) int {
    for i := 2; i <= count; i++ {
    initialfactorial *= i
    }
return initialfactorial
}
