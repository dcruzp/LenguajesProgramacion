# Seminario 4 (Go)



#### 1 - Cual es la forma de procesamiento del código fuente utilizada. (los 3 )

​	El compilador de Go se divide lógicamente en cuatro etapas 

 1. **Análisis léxico y gramatical**

    Análisis del archivo fuente del código. Se convierte la secuencia de cadenas en el archivo en una secuencia `Tokens` para el posterior análisis.(este análisis léxico lo realiza el lexer)

    Al análisis gramatical entra la secuencia de tokens que genera le analizador léxico. Estas secuencias serán analizadas por el analizador gramatical en orden. El procesos de análisis gramatical es seguir la gramática definida por el análisis léxico de abajo hacia arriba. O la especificación de arriba hacia abajo, cada archivo de código fuente de Go se resumira en una estructura `SourceFile`:

     ```
     SourceFile = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" }
     ```

    El analizador sintáctico estándar de Goland  utiliza la gramática de LALR. El resultado del análisis sintáctico es en realidad el árbol de sintaxis abstracta (AST). Cada AST corresponde aun archivo de Go independiente. Este árbol de sintaxis abstracta incluye el nombre del paquete, las constantes definidas, las estructuras y las funciones del archivo actual. 

    Si se produce algún error de sintaxis durante el proceso de análisis, el analizador lo encontrara y el mensaje se imprimirá en la salida estándar. Todo el proceso de compilación también se abortara cuando se produzca un error.

    

 2. **Verificación de Tipos y conversión AST** 

    Después de obtener el árbol de sintaxis abstracta (AST) de un conjunto de archivos, el compilador del lenguaje verificara los tipos definidos y usados en el árbol de sintaxis. La verificación de tipos verificara diferentes tipos de nodos en orden: 

     - Constantes, tipos y tipos de funciones 
     - Asignación e inicializacion de variables 
     - El cuerpo de la función y cierre 
     - Tipos de pares clave- valor hash
     - Importe de cuerpo de la función 
     - Declaración externa.   

    Al recorrer cada árbol de nodos abstractos se verifica el tipo de subarbol actual en cada nodo para asegurar que no habrá errores de en el nodo actual. Todos los errores de tipo y desajuste estarán en esta etapa. 

    La etapa de verificación de tipos no solo verifica los nodos de la estructura del árbol, sino que también expande y reescribe algunas funciones integradas. Por ejemplo la palabra clave make sera reemplazada con funciones como makeslice o makechan en esta etapa de acuerdo con la estructura del subárbol. 

    La verificación de tipos no solo verifica el tipo, sino que también reescribe el AST y procesa las palabras claves integradas del lenguaje Go. Por lo tanto este proceso es muy importante en todo el proceso de compilación. Sin este paso , muchas claves no estarán disponibles. 

 3. **Generación SSA general** 

    Cuando se convierte el archivo fuente en un árbol de sintaxis abstracta, se analiza la gramática de todo el árbol de sintaxis y se realiza la verificación de tipos, entonces el código no tiene problemas de incompilacion o errores gramaticales. El compilador convertirá el AST de entrada en código intermedio. 

    El código intermedio del compilador de Go utiliza la función SSA (Formulario de asignación única estática). Usando esta función en el proceso de generación de código intermedio, se puede analizar fácilmente las variables y fragmentos inútiles en el código. 

    Después de la verificación de tipos, una función llamada compileFunctions comenzara a compilar todas las funcioes en todo el proyecto del lenguaje Go. Estas funciones esperaran el consumo de varias corrutinas de trabajo del back-end en una cola de compilación.

 4. **Generación final del código de maquina** 

    El directorio `cmd/compiler/internal` del código fuente del lenguaje Go contiene una gran cantidad de paquetes relacionados con la generación de código maquina. Los diferentes tipos de CPU usan diferentes paquetes para generar  `amd64` `arm` `mips` `mips64` `ppc64` `s390x` `x86` y `wasm`, lo que significa que el lenguaje Go en casi todos los tipos de conjuntos de instrucciones de CPU comunes.

    
    
  ##### Entrada del compilador 

La entrada del compilador del lenguaje de Go es el archivo main.go en el paquete `src/cmd/compile/internal/gc` . Esta función es el programa principal del compilador del lenguaje Go. Esta función primero obtiene la entrada de la linea de comandos (Parámetro) y actualiza las opciones de compilación y configuración y luego comienza a ejecutar la función `parseFile`  para realizar análisis léxico y gramatical en todos los archivos de entrada para obtener el árbol de sintaxis abstracto correspondiente al archivo. 

A continuación, el árbol de sintaxis abstracta se actualizara y compilara en nueve etapas. Como presentamos anteriormente, todo el proceso pasara por tres partes: verificación de tipos , generación de código intermedio SSA y generación de código maquina.





#### 2 - Go es un lenguaje moderno con muchisimas decisiones de diseño intencionales. Que ventejas  y desventajas le da al lenguaje su forma de procesamiento. Tome en cuenta las plataformas sobre las que se usa para elaborar su respuesta.  (David)

Go es un lenguaje que pensó haciendo énfasis en la simplicidad lo que lo hace fácil de aprender. Su sintaxis es pequeña por lo
que no tendrás que pasar años hojeando la documenación de referencia. El manejo de la memoria y la sintaxis es bastante liviana lo que lo hace fácil de usar.Tiene una compuilación rápida lo que mejora la productividad. Tine un rápido código compilado acercándose bastante a C en ese aspecto. Tiene soporte nativo para la concurrencia lo cual permite escribir
código más simple. Es un leguaje de tipado estático con una librería standard bastante consistente y fácil de instalar para el desarrollo haciendo uso de **go install**. Es autodocumentado y bien documentado . Es libre y de código abierto (licencia BSD).



#### 3 - Realice un sumario sobre las características mas interesantes de la sintaxis de Go: (los 3)

- Presente un Hello World (creatividad apreciada) 

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	fmt.Println("Hello, World!!!")
  }
  ```

- 

- 



## For

Go tiene solo una sola estructura para ciclos **for loop**

Un ciclo **for** básico en go tiene tres componentes principales separadas por semicolons(;):

1.  **init statment**:  la declaración init se ejectuta antes de la primera iteración del ciclo
2.  **condition expression**: la condicional es evaluada antes de cada iteración
3.  **post statement**: se ejectuta al final de cada iteración



La declaración init a menudo será una declaración de variable corta(usando **:=**), y las variables declaradas alli son visibles solo en el scope

de esta instrucción. El ciclo for para de iterar una vez que la condición booleana evaluada es falsa. A diferencia de otros lenguajes como C, Java o Javascript aquí no hay paréntesis que rodeen las tress componentes de la instrucción for y las llaves (**{}**) siempre son necesarias.

Ejemplo de instrucción **for** básico en **Go**:



```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += 2*i
	}
	fmt.Println(sum)
}

```



Las instrucciones init y post son opcionales:

```go
package main

import "fmt"

func main() {
	sum := 0
	for ; sum < 50; {
		sum += 5
	}
	fmt.Println(sum)
}
```

For is también el "while" de Go :):

Puedes quitar el semicolon(;) y el "while" está escritor for en Go

Ejemplo:

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

Ciclo infinito:

Si tu omites la condición booleana sería un ciclo infinito. Por lo tanto puedes expresar un ciclo infinito de forma compacta (no es necesario poner true como en otros lenguajes)

Ejemplo:

```go
package main

func main() {
	for {
	}
}
```






- Creación de variables 

- Ciclos ```for```

- Indentacion 

- Condiciones ```if``` con declaración de varaibles en la condición 

- Funciones con múltiples retornos 

- Otros elementos de las sintaxis que consideres relevante  a mostrar




**If** 

Las sentencias if de Go al igual que for no necesita estar entre paréntesis () pero los llaves  {} si son obligatorias



```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

```

**If** con **short statement**

Similar a for, la instrucción if puede comenzar con una instrucción corta para ejecutar antes de la condición. Las variables declaradas por la instrucción

solo son accesible en el scope hasta el final del if 

Por ejemplo:

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

## If and else

Las variables declaradas dentro de un if short statement son también accesibles dentro de cualquiera de los bloques else

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

Ambas llamadas a **pow**  devuelven sus resultados antes que se haga la llamada a **fmt.Println** en el **main**



### Switch:

Una sentencia **switch** es una manera más corta de escribir una secuencia de  sentencias **if-else**.  Esta ejecuta el primer caso cuyo valor es igual a la expresión de condición .



#### 4 - Presente los tipos nativos  (Daniel)

 1. **Enteros** 

    Go tiene los siguientes tipos de enteros independientes de la arquitectura

    | Nombre | signo    | bits   | type     | Range                                        |
    | ------ | -------- | ------ | -------- | -------------------------------------------- |
    | uint8  | unsigned | 8-bit  | integers | (0 a 255)                                    |
    | uint16 | unsigned | 16-bit | integers | (0 a 65535)                                  |
    | uint32 | unsigned | 32-bit | integers | (0 a 4294967295)                             |
    | uint64 | unsigned | 64-bit | integers | (0 a 18446744073709551615)                   |
    | int8   | signed   | 8-bits | integers | (-128 a 127)                                 |
    | int16  | signed   | 16-bit | integers | (-32768 a 32767)                             |
    | int32  | signed   | 32-bit | integers | (-2147483648 a 2147483647)                   |
    | int64  | signed   | 64-bit | integers | (-9223372036854775808 a 9223372036854775807) |

2. **Flotantes**

   Los números flotantes y complejos también vienen en diferentes tamaños:

   | Nombre  | Signo    | bits   | Type                   |
   | ------- | -------- | ------ | ---------------------- |
   | float32 | IEEE-754 | 32-bit | floating-point numbers |
   | float64 | IEEE-754 | 64-bit | floating-point numbers |

3. **Complejos**

   | Nombre     | Especificaciones                                      |
   | ---------- | ----------------------------------------------------- |
   | complex64  | complex numbers with float32 real and imaginary parts |
   | complex128 | complex numbres with float64 real and imaginary parts |

    También existen varios alias de tipos de números, que asignan nombres útiles a tipos de datos específicos:
   
   | Alias | Tipo de datos |
   | ----- | ------------- |
   | byte  | uint8         |
   | rune  | int32         |

4. **Booleanos**
    El tipo de datos booleanos puede ser uno de los dos valores, ya sea `true` o `false` y se define como `bool` al declararlo como un tipo de datos. Estos valores siempre aparecen con `t` y `f` ya que son identificadores declarado previamente en Go.

5. **Cadenas**

    Una cadena es una secuencia de uno o mas caracteres (letras, números, símbolos) que pueden ser una constante o una variable. Las cadenas existen dentro de comillas invertidas en Go y tienen diferentes características según se utilice.

    Si se utiliza comilla invertida, creara un literal de cadena sin formato. Si utiliza comillas inversas, a veces se conocen como tildes inversas. Dentro de las comillas, cualquier carácter aparecerá como se muestra entre las comillas inversas,  excepción del propio carácter de comilla inversa.

    ```go
    fmt.Println(`Say "Hello World" to Go`)
    ```

    **Output:** ```Say "Hello World" to Go``` 

    Normalmente las barras diagonales inversas se usan para representar caracteres especiales de cadenas. Sin embargo, las barras diagonales inversas no tiene un significado especial dentro de las  caracteres de cadena sin formato. 

    ```go
    fmt.Println(`Say "Hello World" to Go\n`)
    ```

    **Output:** ```Say "Hello World" to Go\n```

  

#### 5 - Comente sobre ```nil``` y los valores por defecto (Daniel) 

	##### Defaults Values 

 1. ##### Valores por defecto para los `string` 

    Los valores por defecto para los string en Go es un string vacio `""`

    ```go
    package main
    
    import "fmt"
    
    func main (){
        var message string 
        fmt.Println(message) 
    }
    ```

    **Output**: ` `

    

	2. ##### Valores por defecto para los enteros 

    En Go existen varios tipos para los enteros primitivos, estos son: `int` `int8` `int16` `int32` `int64` `uint` `uint8` `uint16` `uint32` `uint64` `uintptr` `byte` `rune` . 

     Los valores por defecto para todos esos tipos de enteros en Go son el numero `0` .

    ```go
    package main
    
    import "fmt"
    
    func main() {
       var num int
       var num1 int8
       var num2 uint16
       var num3 uintptr
       var num4 byte
       var num5 rune
    
       fmt.Println(num, num1, num2, num3, num4, num5) // 0, 0, 0, 0, 0, 0
    }
    ```

    **Output**: `0 0 0 0 0 0`

    

	3. ##### Valores por defecto para los Floats

    Como en los enteros en Go hay multiples tipos para representar los floats. Los tipos para los floats son `float32` `float64`

    El valor por defecto para los floats en Go es `0`  , como en los enteros.

    ```go
    package main 
    
    import "fmt" 
    
    func main(){
        var num float32
        var num1 float64
        
        fmt.Println(num,num1)
    }
    ```

    **Output**: `0 0`

	4. ##### Valores por defecto para los Complejos

    En Go tambien existen dos tipos primitivos para representar los numeros complejos que pueden tener o no parte imaginaria. Los dos tipos son `complex64` `complex128` y susu valores por defecto son `(0+0i)`.

    ```go
    package main 
    
    import "fmt" 
    
    func main(){
        var num complex64
        var num1 complex128
        
        fmt.Println(num,num1)
    }
    ```

    **Output**: `(0+0i) (0+0i)`

	5. ##### Valores por defecto para los booleanos

    El valor por defecto para el tipo `bool` en Go es `false`

     

    ```go
    package main
    
    import "fmt"
    
    func main() {
       var facts bool
    
       fmt.Println(facts) // false
    }
    ```

    **Output**: `false`

    

#### 6 - Arrays y slices en Go. Métodos nativos ```make```  ```append```  ``` copy``` Son los slice listas dinámicas? (Javier) 

#### 7 - Los tipos en Go son por referencia o por valor. Punteros en Go, que son , que se puede hacer con ellos , son seguros? Por que es seguro referenciar en Go a una variable local de un metodo?  Haga una comparacion con los punteros de C o C++. (Javier)



#### 8 - Que es el keyword ```defer``` ?  (Daniel)



#### 9 - Presente los ```structs``` en Go y comparelos con los de C. (David)



#### 10.  Que es la composición de tipos? Que son las interfaces en Go? Haga una comparación entre composición de tipos y herencia. Valore ventejas y desventajas de la composición de tipos de Go y exprese su preferencia. (David)



#### 11 - Se puede decir que Go es un lenguaje que ofrece programación orientada a objetos? 



#### 12 - Implemente una jerarquía de clases del seminario de genericidad (Seminario 3) usando ```structs``` e ```interfaces``` .Trate de que los métodos solo reciban tipos nativos o interfaces . Les resulto mas cómodo que usar herencia? Les resulta mas seguro? Les resulta mas expresivo? 

#### 13 - Argumente  el poder que tiene la programación con interfaces para el desarrollo de software, sobre todo el poder que ofrecen las interfaces de Go y C#. 



#### 14 Como maneja Go las excepciones y errores en ejecución? 

#### 15 - Go no presente genericidad de tipos Que limitaciones les puede ofrecer esto al lenguaje? que alternativa propone? 



[TOC]



