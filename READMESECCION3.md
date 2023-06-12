# Primer hola mundo
Para crear archivo compilado
go build main.go
./main

lo que vamos a utilizar es 
go run

si tenemos dos main, podemos ejecutar cada uno, pero el linter nos mandara error, no deben hacer dos main

# Variables

para iniciar un modulo, usamos 
go mod init variables

go mod tidy

luego ya para instalar las dependencias

go get rsc.io/quote

# Declarar constantes

Valores que no cambian

en vez de var es const, esto se define a nivel de paquete o de funcion

# Tipos de datos basicos

# Conversiones de tipos

# El paquete fmt

Este paquete es mas importante de lo que parece, nos ayuda es para manejar valores en la consola

# Operadores aritmeticos y Paquete Math

# Resumen
En la sección de uso de variables y datos en Go, se cubrieron varios temas fundamentales para trabajar con variables y datos en este lenguaje de programación.

En primer lugar, se explicó cómo declarar variables y constantes en Go y cómo asignarles valores. Se presentaron los tipos de datos básicos, incluyendo enteros, flotantes, booleanos y cadenas, y se explicó cómo se utilizan en Go.

También se habló de los valores predeterminados que se asignan a las variables y constantes cuando se declaran sin un valor explícito, y se cubrieron los diferentes métodos para convertir entre diferentes tipos de datos.

Además, se presentó el paquete fmt, que proporciona funciones para imprimir y leer datos de entrada/salida, y se describieron los operadores aritméticos básicos y el paquete math para realizar operaciones matemáticas más complejas.

Por último, se proporcionó un ejemplo práctico en el que se utilizaron los conceptos anteriores para calcular e imprimir el área y el perímetro de un triángulo.

Al comprender estos conceptos, los programadores pueden comenzar a escribir programas simples y comprender mejor cómo trabajar con variables y datos en Go.



Materia de esta sección en GitHub y Tambien puedes descargar en PDF.
https://github.com/alexroel/curso-golang/blob/main/sections/02-variables-data.md


Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

Uso de paquetes, variables y funciones en Go
https://learn.microsoft.com/es-es/training/modules/go-variables-functions-packages/
