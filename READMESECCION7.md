# Iniciar un modulo que otros puedan utilizar

IMPORTANTE
cd espacio de trabajo
mkdir greetings
cd greetings
go mod init github.com/dariel/mygolang
(ya aqui dentro hacemos nuestro codigo y ya podemos llamarlo desde otros modulos)


go build main.go
./main

go build -o hello main.go
./hello

source .profile
hello y se ejecuta

### Desplegar a github
Un despliegue comun y corriente

Resumen
En la sección "Crear un módulo" del curso de Go, aprendimos cómo iniciar un módulo en Go para que otros desarrolladores pudieran utilizarlo en sus proyectos. También exploramos cómo llamar a nuestro código desde otro módulo y cómo manejar y devolver errores adecuadamente.

Además, descubrimos cómo agregar una divertida funcionalidad a nuestro módulo: devolver un saludo aleatorio. Esto le dio un toque especial a nuestras aplicaciones.

Por supuesto, no nos olvidamos de la importancia de las pruebas en el desarrollo de software. Aprendimos cómo agregar pruebas a nuestro módulo para asegurarnos de que funcionaba correctamente en diferentes escenarios.

Una vez que construimos nuestro módulo y probamos su funcionalidad, aprendimos cómo compilar e instalar nuestra aplicación. Esto nos permitió utilizar nuestro módulo en cualquier proyecto sin problemas.

Finalmente, descubrimos cómo desplegar nuestro módulo en GitHub, una plataforma popular para compartir y colaborar en proyectos de código abierto.

Exploramos el emocionante mundo de la creación de módulos en Go y compartimos nuestro código con otros desarrolladores.