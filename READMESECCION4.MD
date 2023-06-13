# Operadores booleanos
Los operadores relacionales y lógicos son utilizados en conjunto en las expresiones lógicas de Go para evaluar condiciones complejas y producir un resultado booleano (verdadero o falso).

Operadores de comparación
Se usan para comparar dos valores y devolver un valor booleano (true o false) según el resultado de la comparación. Los operadores de comparación incluyen:

Igualdad (==)

Desigualdad (!=)

Mayor que (>)

Menor que (<)

Mayor o igual que (>=)

Menor o igual que (<=)

Aquí hay algunos ejemplos de cómo podríamos utilizar los operadores de comparación en Go:

    // Comparación de números
    fmt.Println(1 == 1)   // true
    fmt.Println(1 != 2)   // true
    fmt.Println(2 < 3)    // true
    fmt.Println(3 > 4)    // false
    fmt.Println(4 <= 4)   // true
    fmt.Println(5 >= 6)   // false
 
    // Comparación de cadenas
    fmt.Println("hola" == "hola")       // true
    fmt.Println("hola" != "adios")      // true
    fmt.Println("abc" < "def")          // true
    fmt.Println("ghi" > "fgh")          // true
    fmt.Println("hij" <= "hij")         // true
    fmt.Println("klm" >= "klmno")       // false
 
    // Comparación de booleanos
    fmt.Println(true == true)           // true
    fmt.Println(false != true)          // true
    fmt.Println(true && false == false) // true
    fmt.Println(true || false == true)  // true
En este ejemplo, utilizamos los operadores de comparación para comparar números, cadenas y booleanos. En cada caso, el resultado de la comparación es un valor booleano que representa si la comparación es verdadera o falsa. Los valores booleanos también se pueden comparar utilizando los mismos operadores de comparación.

Operadores lógicos
Los operadores lógicos en Go son utilizados para evaluar expresiones lógicas y producir un resultado booleano (verdadero o falso). En Go, existen tres operadores lógicos: AND lógico (&&), OR lógico (||) y NOT lógico (!).

Operador AND lógico (&&):
El operador && evalúa dos expresiones booleanas y devuelve verdadero (true) si ambas expresiones son verdaderas, y devuelve falso (false) si alguna de las expresiones es falsa. Por ejemplo:

x := true
y := false
z := x && y
fmt.Println(z) // Imprime false
En este ejemplo, la variable z tendrá el valor de falso (false), ya que la expresión x && y evalúa a falso, debido a que la variable y es falsa.

Operador OR lógico (||):
El operador || evalúa dos expresiones booleanas y devuelve verdadero (true) si al menos una de las expresiones es verdadera, y devuelve falso (false) si ambas expresiones son falsas. Por ejemplo:

x := true
y := false
z := x || y
fmt.Println(z) // Imprime true
En este ejemplo, la variable z tendrá el valor de verdadero (true), ya que la expresión x || y evalúa a verdadero, debido a que la variable x es verdadera.

Operador NOT lógico (!):
El operador ! niega el valor booleano de una expresión, es decir, si una expresión es verdadera, la niega y devuelve falso, y si una expresión es falsa, la niega y devuelve verdadero. Por ejemplo:

x := true
z := !x
fmt.Println(z) // Imprime false
En este ejemplo, la variable z tendrá el valor de falso (false), ya que la expresión !x niega el valor de x, que es verdadero, y devuelve falso.

Aquí hay un ejemplo en Go que utiliza solo operadores lógicos para realizar operaciones sin condiciones:

    x := true
    y := false
 
    // Negación
    fmt.Println(!x) // false
    fmt.Println(!y) // true
 
    // AND lógico
    fmt.Println(x && x) // true
    fmt.Println(x && y) // false
    fmt.Println(y && y) // false
 
    // OR lógico
    fmt.Println(x || x) // true
    fmt.Println(x || y) // true
    fmt.Println(y || y) // false
En este ejemplo, utilizamos los operadores lógicos !, && y || para realizar operaciones sin utilizar condiciones explícitas.

En la primera sección, utilizamos el operador de negación ! para invertir los valores de x e y.

En la sección de operaciones lógicas && y ||, realizamos operaciones booleanas utilizando los valores de x e y. En el caso del operador &&, el resultado será verdadero solo si ambos operandos son verdaderos, en otro caso, el resultado es falso. En el caso del operador ||, el resultado será verdadero si al menos uno de los operandos es verdadero, de lo contrario, el resultado será falso.

Expresiones
En programación, una expresión es una combinación de valores, variables, operadores y llamadas a funciones que se evalúa para producir un resultado. Las expresiones pueden ser tan simples como una variable que se asigna a un valor, o tan complejas como una ecuación matemática con múltiples operadores y variables.

El orden en que se resuelven las expresiones en un programa depende de la prioridad de los operadores y los paréntesis utilizados para agrupar las operaciones. El orden de resolución de las operaciones sigue las reglas matemáticas convencionales:

Los paréntesis se evalúan primero. Las expresiones dentro de los paréntesis se resuelven antes que cualquier otra operación.

Luego se resuelven las operaciones aritméticas, como la multiplicación, la división, la suma y la resta. La multiplicación y la división se resuelven antes que la suma y la resta.

Finalmente, se resuelven las operaciones de comparación y los operadores lógicos.

Es importante tener en cuenta que los operadores con la misma prioridad se resuelven de izquierda a derecha. Por ejemplo, en la expresión 2 + 3 * 4, la multiplicación se resuelve primero debido a su mayor prioridad, y el resultado es 14. Si queremos que la suma se resuelva primero, debemos utilizar paréntesis para indicar la prioridad, como en (2 + 3) * 4, lo que da como resultado 20.

El orden de resolución de las expresiones es importante porque puede afectar el resultado final de un programa. Por lo tanto, es importante comprender la prioridad de los operadores y utilizar paréntesis para agrupar las operaciones de la manera adecuada.

Aquí hay un ejemplo en Go que utiliza expresiones con paréntesis, operadores aritméticos, operadores de comparación y operadores lógicos.

    x := 5
    y := 10
    z := 15
 
    // Expresión con paréntesis, operadores aritméticos, de comparación y lógicos
    resultado := ((x+y)*z)/(x*y) > z && x != y
 
    // Imprimir el resultado
    fmt.Println(resultado) //False
En este ejemplo, definimos tres variables x, y y z con valores enteros. Luego, utilizamos estos valores para construir una expresión que incluye paréntesis, operadores aritméticos, operadores de comparación y operadores lógicos.

En lugar de utilizar la estructura de control if para evaluar la expresión, simplemente asignamos el resultado de la expresión a una variable llamada resultado. La expresión se evalúa de la misma manera que en el ejemplo anterior.

Finalmente, imprimimos el valor de la variable resultado utilizando la función fmt.Println(). Si la expresión se evalúa como verdadera, se imprimirá true. Si se evalúa como falsa, se imprimirá false.

# Declaracion If - Else

# Bucle For

# Hacer nuestra propia funcion

# Proyecto de la seccion, numero aleatorio de cero a cien

# Resumen

En esta sección del curso de programación en Go, hemos explorado el control de flujo. Hemos cubierto varios temas importantes, que incluyen:

Operadores booleanos: hemos aprendido acerca de los operadores booleanos en Go y cómo utilizarlos para tomar decisiones en nuestro código.

Declaración If-Else: hemos aprendido cómo utilizar la declaración if-else en Go para tomar decisiones en función de las condiciones en nuestro programa.

Declaración Switch: hemos aprendido cómo utilizar la declaración switch en Go para tomar decisiones en función de múltiples condiciones.

Bucle For: hemos aprendido cómo utilizar el bucle for en Go para repetir una sección de código un número determinado de veces.

Uso de funciones: hemos aprendido cómo utilizar funciones en Go y cómo crear nuestras propias funciones personalizadas.

Además, al final de la sección, hemos tenido la oportunidad de aplicar todo lo que hemos aprendido en un proyecto práctico, donde hemos creado un programa que utiliza los conceptos de control de flujo que hemos cubierto.

En general, esta sección ha sido una introducción muy útil al control de flujo en Go, y proporciona una base sólida para continuar aprendiendo y desarrollando habilidades en este lenguaje de programación.



Materia de esta sección en GitHub y Tambien puedes descargar en PDF.



Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

https://learn.microsoft.com/es-es/training/modules/go-control-flow/