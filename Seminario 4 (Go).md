# Seminario 4 (Go)



#### 1 - Cual es la forma de procesamiento del código fuente utilizada.

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





#### 2 - Go es un lenguaje moderno con muchisimas decisiones de diseño intencionales. Que ventejas  y desventajas le da al lenguaje su forma de procesamiento. Tome en cuenta las plataformas sobre las que se usa para elaborar su respuesta.  



#### 3 - Realice un sumario sobre las características mas interesantes de la sintaxis de Go:

- Presente un Hello World (creatividad apreciada) 

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	fmt.Println("Hello, World!")
  }
  ```

  

- Creación de variables 

- Ciclos ```for```

- Indentacion 

- Condiciones ```if``` con declaración de varaibles en la condición 

- Funciones con múltiples retornos 

- Otros elementos de las sintaxis que consideres relevante  a mostrar

  
#### 4 - Presente los tipos nativos

#### 5 - Comente sobre ```nil``` y los valores por defecto 

#### 6 - Arrays y slices en Go. Métodos nativos ```make```  ```append```  ``` copy``` Son los slice listas dinámicas?

#### 7 - Los tipos en Go son por referencia o por valor. Punteros en Go, que son , que se puede hacer con ellos , son seguros? Por que es seguro referenciar en Go a una variable local de un metodo?  Haga una comparacion con los punteros de C o C++.



#### 8 - Que es el keyword ```defer``` ? 



#### 9 - Presente los ```structs``` en Go y comparelos con los de C. 



#### 10.  Que es la composicion de tipos? Que son las interfaces en Go? Haga una comparacion entre composicion de tipos y herencia. Valore ventejas y desventajas de la composicion de tipos de Go y exprese su preferencia. 



#### 11 - Se puede decir que Go es un lenguaje que ofrece programacion orientada a objetos? 



#### 12 - Implemente una gerarquia de clases del seminario de genericidad (Seminario 3) usando ```structs``` e ```interfaces``` .Trate de que los metodos solo reciban tipos nativos o interfaces . Les resulto mas comodo que usar herencia? Les resulta mas seguro? Les resulta mas expresivo? 

#### 13 - Argumente  el poder que tiene la programacion con interfaces para el desarrollo de software, sobre todo el poder que ofrecen las interfaces de Go y C#. 



#### 14 Como maneja Go las excepciones y errores en ejecucion? 

#### 15 - Go no presente genericidad de tipos Que limitaciones les puede ofrecer esto el lenguaje? que alternativa propone? 



[TOC]



