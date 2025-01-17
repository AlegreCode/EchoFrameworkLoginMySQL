# EchoFrameworkLoginMySQL
### Plantilla de inicio de sesión y registro de usuario con Echo framework y MySQL.

##### Packages utilizados:
- **Sprig**: extiende las funcionalidades de template de golang. [Aquí](https://github.com/Masterminds/sprig)
- **gookit/validate**: agrega funciones de validación de entradas. [Aquí](https://github.com/gookit/validate)
- **godotenv**: permite cargar variables de entorno desde un archivo .env. [Aquí](https://github.com/joho/godotenv)
- **bcrypt**: utilizado para encriptar las contraseñas de los usuarios. [Aquí](https://godoc.org/golang.org/x/crypto/bcrypt)
- **go-sql-driver/mysql**: agrega el driver para conexiones mysql. [Aquí](http://go-database-sql.org/index.html)

##### Config base de datos
El archivo .env-example renombrar a .env y completar las variables con tus credenciales de conexión.

##### Instalación de paquetes
Para evitar conflictos de versiones utilizamos el gestor de dependencias de go [dep](https://golang.github.io/dep/). Debes tener instalado esta herramienta (para ver como clic [aquí](https://golang.github.io/dep/docs/installation.html)), luego entrar a la raíz de tu proyecto y ejecutar el comando `dep ensure`. Una vez instalados todos los paquetes ya correr el proyecto.
