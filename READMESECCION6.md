Manejo de errores
Uso de defer
Uso de panic y recover
Registro de errores
Proyecto: Gestor de contacto

# Manejo de errores
Garantizar la robustes y la calidad

# Uso de defer
Defer nos sirve para posponer la ejecucion de una funcion hasta que la funcion que la contiene haya finalizado
Esto es muy bueno para realizar cosas cuando se termina una ejecucion

# Uso de panic y recover
panic y recover se utilizan para manejar errores graves o situaciones excepcionales
recover muestra el panico, sin interrumpir el programa
No se recomienda para manejar errores, solo para situaciones de limpieza y recuperacion, o excepcionales

# Registro de errores
Como desarrollador necesitamos mas informacion para hacer una correccion, para registrar errores ver ejemplo


# Resumen
En la sección del curso de Go dedicada al manejo de errores, exploramos diferentes aspectos del manejo de errores en Go y cómo aplicarlos de manera efectiva en nuestras aplicaciones.

Comenzamos aprendiendo sobre el manejo de errores en Go, donde vimos cómo utilizar la sintaxis if err != nil para verificar y manejar errores de forma concisa y eficiente. También exploramos las mejores prácticas para devolver y manejar errores adecuadamente en nuestras funciones.

Luego profundizamos en un ejemplo práctico de manejo de errores en Go, donde analizamos cómo implementar el manejo de errores en una aplicación real. Exploramos técnicas como el uso de variables de error, el registro de errores y la propagación adecuada de errores en diferentes capas de nuestra aplicación.

Uno de los aspectos más interesantes del manejo de errores en Go es el uso de la declaración defer. Aprendimos cómo usar defer para asegurarnos de que ciertas operaciones se realicen antes de salir de una función, incluso en caso de error. Vimos ejemplos prácticos que demostraron cómo defer puede facilitar el manejo de tareas como la liberación de recursos o el cierre adecuado de archivos.

También exploramos el uso de panic y recover. Aprendimos cómo utilizar estas dos funciones para manejar situaciones excepcionales y errores graves en nuestras aplicaciones. Vimos cómo panic puede interrumpir el flujo normal del programa y cómo recover nos permite recuperarnos de dichos errores y continuar la ejecución de manera controlada.

En el proyecto de esta sección, creamos un gestor de contactos utilizando Go. Aplicamos todos los conceptos que aprendimos sobre el manejo de errores en Go para implementar una funcionalidad robusta para agregar, mostrar y almacenar contactos, asegurándonos de manejar los posibles errores que pudieran ocurrir durante el proceso.

Concluimos esta sección con una sólida comprensión del manejo de errores en Go. Ahora estamos preparados para crear aplicaciones confiables y resistentes que manejen adecuadamente situaciones inesperadas.



Materia de esta sección en GitHub y Tambien puedes descargar en PDF.
https://github.com/alexroel/curso-golang/blob/main/sections/05-error-handling.md


Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

Implementación del control de errores y el registro en Go
https://learn.microsoft.com/es-es/training/modules/go-errors-logs/